package main

import (
	log "code.google.com/p/log4go"
	"errors"
)

var (
	ErrRingEmpty = errors.New("ring buffer empty")
	ErrRingFull  = errors.New("ring buffer full")
)

type Ring struct {
	// read
	rn uint
	rp int
	// write TODO cpu cache line aligned?
	wn uint
	wp int
	// info
	num  int
	data []Proto
}

func NewRing(num int) *Ring {
	r := new(Ring)
	r.num = num
	r.data = make([]Proto, num)
	return r
}

func InitRing(r *Ring, num int) {
	r.num = num
	r.data = make([]Proto, num)
}

func (r *Ring) Get() (proto *Proto, err error) {
	if r.wn-r.rn == 0 {
		return nil, ErrRingEmpty
	}
	proto = &r.data[r.rp]
	return
}

func (r *Ring) GetAdv() {
	if r.rp++; r.rp >= r.num {
		r.rp = 0
	}
	r.rn++
	log.Debug("ring rn: %d, rp: %d", r.rn, r.rp)
}

func (r *Ring) Set() (proto *Proto, err error) {
	if r.wn-r.rn >= uint(r.num) {
		return nil, ErrRingFull
	}
	proto = &r.data[r.wp]
	return
}

func (r *Ring) SetAdv() {
	if r.wp++; r.wp >= r.num {
		r.wp = 0
	}
	r.wn++
	log.Debug("ring wn: %d, wp: %d", r.wn, r.wp)
}

func (r *Ring) Reset() {
	r.rn = 0
	r.rp = 0
	r.wn = 0
	r.wp = 0
}

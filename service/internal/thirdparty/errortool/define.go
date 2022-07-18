package errortool

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"
)

const (
	groupCodeDB int = iota
)

const (
	groupPrefixIDMaxSequence       = 1295
	groupPrefixSystemIDMaxSequence = 100 // 保留 1~100 給 default 使用
	groupNumIDMaxSequence          = 46655
)

func Define() *define {
	return &define{
		seq:   newSequence(groupPrefixSystemIDMaxSequence, groupPrefixIDMaxSequence),
		codes: newCodeRepository(),
	}
}

type define struct {
	seq   *sequence
	codes iCodeRepository
}

func (d *define) Group() *errorGroup {
	return &errorGroup{
		codes:     d.codes,
		groupCode: d.seq.Next(),
		seq:       newSequence(0, groupNumIDMaxSequence),
	}
}

func (d *define) defaultGroup(groupCode int) *errorGroup {
	return &errorGroup{
		codes:     d.codes,
		groupCode: groupCode,
		seq:       newSequence(0, groupNumIDMaxSequence),
	}
}

func (d *define) Plugin(f func(codes iCodeRepository) interface{}) interface{} {
	return f(d.codes)
}

func (d *define) List() []errorString {
	keys := d.codes.Keys()
	sort.SliceStable(keys,
		func(i, j int) bool {
			return keys[i] < keys[j]
		})

	res := make([]errorString, len(keys))
	for i, v := range keys {
		if val, ok := d.codes.Get(v); ok {
			res[i] = *val
		} else {
			res[i] = errorString{}
		}
	}

	return res
}

type errorGroup struct {
	codes     iCodeRepository
	groupCode int
	seq       *sequence
}

func (e *errorGroup) Error(message string) error {
	code := e.makeErrorCode(e.groupCode, e.seq.Next())
	err := &errorString{
		code:    code,
		message: message,
	}
	e.codes.Add(code, err)
	return err
}

func (e *errorGroup) makeErrorCode(groupCode, code int) errorCode {
	return errorCode(fmt.Sprintf("%02s-%03d", strings.ToUpper(strconv.FormatInt(int64(groupCode), 36)), code))
}

func newSequence(begin int, max int) *sequence {
	return &sequence{
		now: begin,
		max: max,
	}
}

type sequence struct {
	mx  sync.Mutex
	now int
	max int
}

func (s *sequence) Next() int {
	s.mx.Lock()
	defer s.mx.Unlock()

	s.now++
	if s.now > s.max {
		panic("max sequence")
	}
	return s.now
}

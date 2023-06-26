// Copyright 2018 The gVisor Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package stack

import (
	"github.com/sagernet/gvisor/pkg/tcpip"
	"golang.org/x/time/rate"
)

const (
	// icmpLimit is the default maximum number of ICMP messages permitted by this
	// rate limiter.
	icmpLimit = 1000

	// icmpBurst is the default number of ICMP messages that can be sent in a single
	// burst.
	icmpBurst = 50
)

// ICMPRateLimiter is a global rate limiter that controls the generation of
// ICMP messages generated by the stack.
type ICMPRateLimiter struct {
	limiter *rate.Limiter
	clock   tcpip.Clock
}

// NewICMPRateLimiter returns a global rate limiter for controlling the rate
// at which ICMP messages are generated by the stack. The returned limiter
// does not apply limits to any ICMP types by default.
func NewICMPRateLimiter(clock tcpip.Clock) *ICMPRateLimiter {
	return &ICMPRateLimiter{
		clock:   clock,
		limiter: rate.NewLimiter(icmpLimit, icmpBurst),
	}
}

// SetLimit sets a new Limit for the limiter.
func (l *ICMPRateLimiter) SetLimit(limit rate.Limit) {
	l.limiter.SetLimitAt(l.clock.Now(), limit)
}

// Limit returns the maximum overall event rate.
func (l *ICMPRateLimiter) Limit() rate.Limit {
	return l.limiter.Limit()
}

// SetBurst sets a new burst size for the limiter.
func (l *ICMPRateLimiter) SetBurst(burst int) {
	l.limiter.SetBurstAt(l.clock.Now(), burst)
}

// Burst returns the maximum burst size.
func (l *ICMPRateLimiter) Burst() int {
	return l.limiter.Burst()
}

// Allow reports whether one ICMP message may be sent now.
func (l *ICMPRateLimiter) Allow() bool {
	return l.limiter.AllowN(l.clock.Now(), 1)
}

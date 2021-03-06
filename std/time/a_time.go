// This file is generated by generate-std.joke script. Do not edit manually!

package time

import (
	. "github.com/candid82/joker/core"
	"time"
)

var timeNamespace = GLOBAL_ENV.EnsureNamespace(MakeSymbol("joker.time"))

var ansi_c_ String
var hour_ *BigInt
var kitchen_ String
var microsecond_ Int
var millisecond_ Int
var minute_ *BigInt
var nanosecond_ Int
var rfc1123_ String
var rfc1123_z_ String
var rfc3339_ String
var rfc3339_nano_ String
var rfc822_ String
var rfc822_z_ String
var rfc850_ String
var ruby_date_ String
var second_ Int
var stamp_ String
var stamp_micro_ String
var stamp_milli_ String
var stamp_nano_ String
var unix_date_ String

var add_ Proc

func __add_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 2:
		t := ExtractTime(_args, 0)
		d := ExtractInt(_args, 1)
		_res := t.Add(time.Duration(d))
		return MakeTime(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var add_date_ Proc

func __add_date_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 4:
		t := ExtractTime(_args, 0)
		years := ExtractInt(_args, 1)
		months := ExtractInt(_args, 2)
		days := ExtractInt(_args, 3)
		_res := t.AddDate(years, months, days)
		return MakeTime(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var format_ Proc

func __format_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 2:
		t := ExtractTime(_args, 0)
		layout := ExtractString(_args, 1)
		_res := t.Format(layout)
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var from_unix_ Proc

func __from_unix_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 2:
		sec := ExtractInt(_args, 0)
		nsec := ExtractInt(_args, 1)
		_res := time.Unix(int64(sec), int64(nsec))
		return MakeTime(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var hours_ Proc

func __hours_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		d := ExtractInt(_args, 0)
		_res := time.Duration(d).Hours()
		return MakeDouble(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var minutes_ Proc

func __minutes_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		d := ExtractInt(_args, 0)
		_res := time.Duration(d).Minutes()
		return MakeDouble(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var now_ Proc

func __now_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 0:
		_res := time.Now()
		return MakeTime(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var parse_ Proc

func __parse_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 2:
		layout := ExtractString(_args, 0)
		value := ExtractString(_args, 1)
		_res, err := time.Parse(layout, value)
		PanicOnErr(err)
		return MakeTime(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var parse_duration_ Proc

func __parse_duration_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		s := ExtractString(_args, 0)
		t, err := time.ParseDuration(s)
		PanicOnErr(err)
		_res := int(t)
		return MakeInt(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var round_ Proc

func __round_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 2:
		d := ExtractInt(_args, 0)
		m := ExtractInt(_args, 1)
		_res := int(time.Duration(d).Round(time.Duration(m)))
		return MakeInt(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var seconds_ Proc

func __seconds_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		d := ExtractInt(_args, 0)
		_res := time.Duration(d).Seconds()
		return MakeDouble(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var since_ Proc

func __since_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		t := ExtractTime(_args, 0)
		_res := int(time.Since(t))
		return MakeInt(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var sleep_ Proc

func __sleep_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		d := ExtractInt(_args, 0)
		 RT.GIL.Unlock()
		time.Sleep(time.Duration(d))
		RT.GIL.Lock()
		_res := NIL
		return _res

	default:
		PanicArity(_c)
	}
	return NIL
}

var string_ Proc

func __string_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		d := ExtractInt(_args, 0)
		_res := time.Duration(d).String()
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var sub_ Proc

func __sub_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 2:
		t := ExtractTime(_args, 0)
		u := ExtractTime(_args, 1)
		_res := int(t.Sub(u))
		return MakeInt(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var truncate_ Proc

func __truncate_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 2:
		d := ExtractInt(_args, 0)
		m := ExtractInt(_args, 1)
		_res := int(time.Duration(d).Truncate(time.Duration(m)))
		return MakeInt(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var unix_ Proc

func __unix_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		t := ExtractTime(_args, 0)
		_res := int(t.Unix())
		return MakeInt(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var until_ Proc

func __until_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		t := ExtractTime(_args, 0)
		_res := int(time.Until(t))
		return MakeInt(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

func Init() {
	ansi_c_ = MakeString(time.ANSIC)
	hour_ = MakeBigInt(int64(time.Hour))
	kitchen_ = MakeString(time.Kitchen)
	microsecond_ = MakeInt(int(time.Microsecond))
	millisecond_ = MakeInt(int(time.Millisecond))
	minute_ = MakeBigInt(int64(time.Minute))
	nanosecond_ = MakeInt(int(time.Nanosecond))
	rfc1123_ = MakeString(time.RFC1123)
	rfc1123_z_ = MakeString(time.RFC1123Z)
	rfc3339_ = MakeString(time.RFC3339)
	rfc3339_nano_ = MakeString(time.RFC3339Nano)
	rfc822_ = MakeString(time.RFC822)
	rfc822_z_ = MakeString(time.RFC822Z)
	rfc850_ = MakeString(time.RFC850)
	ruby_date_ = MakeString(time.RubyDate)
	second_ = MakeInt(int(time.Second))
	stamp_ = MakeString(time.Stamp)
	stamp_micro_ = MakeString(time.StampMicro)
	stamp_milli_ = MakeString(time.StampMilli)
	stamp_nano_ = MakeString(time.StampNano)
	unix_date_ = MakeString(time.UnixDate)
	add_ = __add_
	add_date_ = __add_date_
	format_ = __format_
	from_unix_ = __from_unix_
	hours_ = __hours_
	minutes_ = __minutes_
	now_ = __now_
	parse_ = __parse_
	parse_duration_ = __parse_duration_
	round_ = __round_
	seconds_ = __seconds_
	since_ = __since_
	sleep_ = __sleep_
	string_ = __string_
	sub_ = __sub_
	truncate_ = __truncate_
	unix_ = __unix_
	until_ = __until_

	timeNamespace.ResetMeta(MakeMeta(nil, `Provides functionality for measuring and displaying time.`, "1.0"))

	timeNamespace.InternVar("ansi-c", ansi_c_,
		MakeMeta(
			nil,
			`Mon Jan _2 15:04:05 2006`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

	timeNamespace.InternVar("hour", hour_,
		MakeMeta(
			nil,
			`Number of nanoseconds in 1 hour`, "1.0").Plus(MakeKeyword("tag"), String{S: "BigInt"}))

	timeNamespace.InternVar("kitchen", kitchen_,
		MakeMeta(
			nil,
			`3:04PM`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

	timeNamespace.InternVar("microsecond", microsecond_,
		MakeMeta(
			nil,
			`Number of nanoseconds in 1 microsecond`, "1.0").Plus(MakeKeyword("tag"), String{S: "Int"}))

	timeNamespace.InternVar("millisecond", millisecond_,
		MakeMeta(
			nil,
			`Number of nanoseconds in 1 millisecond`, "1.0").Plus(MakeKeyword("tag"), String{S: "Int"}))

	timeNamespace.InternVar("minute", minute_,
		MakeMeta(
			nil,
			`Number of nanoseconds in 1 minute`, "1.0").Plus(MakeKeyword("tag"), String{S: "BigInt"}))

	timeNamespace.InternVar("nanosecond", nanosecond_,
		MakeMeta(
			nil,
			`Number of nanoseconds in 1 nanosecond`, "1.0").Plus(MakeKeyword("tag"), String{S: "Int"}))

	timeNamespace.InternVar("rfc1123", rfc1123_,
		MakeMeta(
			nil,
			`Mon, 02 Jan 2006 15:04:05 MST`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

	timeNamespace.InternVar("rfc1123-z", rfc1123_z_,
		MakeMeta(
			nil,
			`Mon, 02 Jan 2006 15:04:05 -0700`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

	timeNamespace.InternVar("rfc3339", rfc3339_,
		MakeMeta(
			nil,
			`2006-01-02T15:04:05Z07:00`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

	timeNamespace.InternVar("rfc3339-nano", rfc3339_nano_,
		MakeMeta(
			nil,
			`2006-01-02T15:04:05.999999999Z07:00`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

	timeNamespace.InternVar("rfc822", rfc822_,
		MakeMeta(
			nil,
			`02 Jan 06 15:04 MST`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

	timeNamespace.InternVar("rfc822-z", rfc822_z_,
		MakeMeta(
			nil,
			`02 Jan 06 15:04 -0700`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

	timeNamespace.InternVar("rfc850", rfc850_,
		MakeMeta(
			nil,
			`Monday, 02-Jan-06 15:04:05 MST`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

	timeNamespace.InternVar("ruby-date", ruby_date_,
		MakeMeta(
			nil,
			`Mon Jan 02 15:04:05 -0700 2006`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

	timeNamespace.InternVar("second", second_,
		MakeMeta(
			nil,
			`Number of nanoseconds in 1 second`, "1.0").Plus(MakeKeyword("tag"), String{S: "Int"}))

	timeNamespace.InternVar("stamp", stamp_,
		MakeMeta(
			nil,
			`Jan _2 15:04:05`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

	timeNamespace.InternVar("stamp-micro", stamp_micro_,
		MakeMeta(
			nil,
			`Jan _2 15:04:05.000000`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

	timeNamespace.InternVar("stamp-milli", stamp_milli_,
		MakeMeta(
			nil,
			`Jan _2 15:04:05.000`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

	timeNamespace.InternVar("stamp-nano", stamp_nano_,
		MakeMeta(
			nil,
			`Jan _2 15:04:05.000000000`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

	timeNamespace.InternVar("unix-date", unix_date_,
		MakeMeta(
			nil,
			`Mon Jan _2 15:04:05 MST 2006`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

	timeNamespace.InternVar("add", add_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("t"), MakeSymbol("d"))),
			`Returns the time t+d.`, "1.0").Plus(MakeKeyword("tag"), String{S: "Time"}))

	timeNamespace.InternVar("add-date", add_date_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("t"), MakeSymbol("years"), MakeSymbol("months"), MakeSymbol("days"))),
			`Returns the time t + (years, months, days).`, "1.0").Plus(MakeKeyword("tag"), String{S: "Time"}))

	timeNamespace.InternVar("format", format_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("t"), MakeSymbol("layout"))),
			`Returns a textual representation of the time value formatted according to layout,
  which defines the format by showing how the reference time, defined to be
  Mon Jan 2 15:04:05 -0700 MST 2006
  would be displayed if it were the value; it serves as an example of the desired output.
  The same display rules will then be applied to the time value..`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

	timeNamespace.InternVar("from-unix", from_unix_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("sec"), MakeSymbol("nsec"))),
			`Returns the local Time corresponding to the given Unix time, sec seconds and
  nsec nanoseconds since January 1, 1970 UTC. It is valid to pass nsec outside the range [0, 999999999].`, "1.0").Plus(MakeKeyword("tag"), String{S: "Time"}))

	timeNamespace.InternVar("hours", hours_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("d"))),
			`Returns the duration (passed as a number of nanoseconds) as a floating point number of hours.`, "1.0").Plus(MakeKeyword("tag"), String{S: "Double"}))

	timeNamespace.InternVar("minutes", minutes_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("d"))),
			`Returns the duration (passed as a number of nanoseconds) as a floating point number of minutes.`, "1.0").Plus(MakeKeyword("tag"), String{S: "Double"}))

	timeNamespace.InternVar("now", now_,
		MakeMeta(
			NewListFrom(NewVectorFrom()),
			`Returns the current local time.`, "1.0").Plus(MakeKeyword("tag"), String{S: "Time"}))

	timeNamespace.InternVar("parse", parse_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("layout"), MakeSymbol("value"))),
			`Parses a time string.`, "1.0").Plus(MakeKeyword("tag"), String{S: "Time"}))

	timeNamespace.InternVar("parse-duration", parse_duration_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("s"))),
			`Parses a duration string. A duration string is a possibly signed sequence of decimal numbers,
  each with optional fraction and a unit suffix, such as 300ms, -1.5h or 2h45m. Valid time units are
  ns, us (or µs), ms, s, m, h.`, "1.0").Plus(MakeKeyword("tag"), String{S: "Int"}))

	timeNamespace.InternVar("round", round_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("d"), MakeSymbol("m"))),
			`Returns the result of rounding d to the nearest multiple of m. d and m represent time durations in nanoseconds.
  The rounding behavior for halfway values is to round away from zero. If m <= 0, returns d unchanged.`, "1.0").Plus(MakeKeyword("tag"), String{S: "Int"}))

	timeNamespace.InternVar("seconds", seconds_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("d"))),
			`Returns the duration (passed as a number of nanoseconds) as a floating point number of seconds.`, "1.0").Plus(MakeKeyword("tag"), String{S: "Double"}))

	timeNamespace.InternVar("since", since_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("t"))),
			`Returns the time in nanoseconds elapsed since t.`, "1.0").Plus(MakeKeyword("tag"), String{S: "Int"}))

	timeNamespace.InternVar("sleep", sleep_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("d"))),
			`Pauses the execution thread for at least the duration d (expressed in nanoseconds).
  A negative or zero duration causes sleep to return immediately.`, "1.0"))

	timeNamespace.InternVar("string", string_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("d"))),
			`Returns a string representing the duration in the form 72h3m0.5s.`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

	timeNamespace.InternVar("sub", sub_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("t"), MakeSymbol("u"))),
			`Returns the duration t-u in nanoseconds.`, "1.0").Plus(MakeKeyword("tag"), String{S: "Int"}))

	timeNamespace.InternVar("truncate", truncate_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("d"), MakeSymbol("m"))),
			`Returns the result of rounding d toward zero to a multiple of m. If m <= 0, returns d unchanged.`, "1.0").Plus(MakeKeyword("tag"), String{S: "Int"}))

	timeNamespace.InternVar("unix", unix_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("t"))),
			`Returns t as a Unix time, the number of seconds elapsed since January 1, 1970 UTC.`, "1.0").Plus(MakeKeyword("tag"), String{S: "Int"}))

	timeNamespace.InternVar("until", until_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("t"))),
			`Returns the duration in nanoseconds until t.`, "1.0").Plus(MakeKeyword("tag"), String{S: "Int"}))

}

func init() {
	timeNamespace.Lazy = Init
}

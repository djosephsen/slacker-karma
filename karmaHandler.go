package karma

import (
	sl "github.com/djosephsen/slacker/slackerlib"
	"regexp"
	"fmt"
	//"strings"
	"strconv"
)

var OrgTracker = sl.MessageHandler{
	Name: `OrgTracker`,
	Usage:`<botname> [add|delete|join|leave|list] org <orgname>`,
	Method: `HEAR`,
	Pattern: `^(\+|-)(\d+) *(\w+)`,
	Run:	func(e *sl.Event, match []string){
		var karma []byte
		var ikarma,ival,newKarma int64
		var err error
		brain := *e.Sbot.Brain
		mod:=match[1]
		val:=match[2]
		peep:=match[3]

		if hasColon,_ := regexp.MatchString(`:$`,peep); hasColon{
			peep=peep[:len(peep)-1]
		}

		user := e.Sbot.Meta.GetUserByName(peep)
		if user == nil{
			sl.Logger.Debug(`Karma:: %s doesn't look like a valid User. Ignoring`,peep)
			return
		}

		key:=`karma::`+user.ID

		if karma,err = brain.Get(key); err != nil{
			sl.Logger.Debug(err)
			return
		}

		if ikarma,err = strconv.ParseInt(string(karma), 0, 64); err != nil{
			sl.Logger.Debug(err)
			return
		}

		if ival,err = strconv.ParseInt(val, 0, 64); err != nil{
			sl.Logger.Debug(err)
			return
		}

		if mod==`+`{
			newKarma = ikarma+ival
		}else if mod==`-`{
			newKarma = ikarma+ival
		}else{
			sl.Logger.Debug("Karma:: funky operator: %s",mod)
			return
		}

		if err := brain.Set(key,[]byte(strconv.FormatInt(newKarma,10))); err != nil{
			sl.Logger.Debug(err)
			return
		}
		e.Reply(fmt.Sprintf("%s's karma is now: %d",peep,newKarma))
	},
}

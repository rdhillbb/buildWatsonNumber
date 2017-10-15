package buildWatsonNumber
//package main

import "fmt"

var msgPart1 string = "<p><s>Alert, <break time=\"25ms\"/>A P1 Case has been filed by <break time=\"450ms\"/>"
var msgPart2 string = "<break time=\"90ms\"/><prosody rate=\"-89%\">Case Number "
var msgPart3 string = "</prosody>Repeat a P1 Case has been filed by <break time=\"300ms\"/>"
var msgPart4 string = "<break time=\"90ms\"/><prosody rate=\"-85%\">Case Number " 
var msgPart5 string = "</prosody> </s> <s> Please contact the TAC immediately to speak to a duty manager </s> Thank You</p>" 

func createNumbMsg(caseNumber string) string {
	numberwSpace := createMsgString(512)
	for _, c := range caseNumber {
		numberwSpace += string(c) + " "
	}
	return numberwSpace
}
func CrMsg4wat(customer, caseNumber string) string {
	watmsg := createMsgString(2048)
	fmt.Println("MY NEW NUMBER", createNumbMsg(caseNumber))
	fmt.Println(msgPart1+customer+msgPart2+createNumbMsg(caseNumber)+msgPart3+customer+msgPart4+createNumbMsg(caseNumber)+msgPart5)
	watmsg =msgPart1+customer+msgPart2+createNumbMsg(caseNumber)+msgPart3+customer+msgPart4+createNumbMsg(caseNumber)+msgPart5
	return watmsg
}
func createMsgString(sizeMsg int) string {
	b := make([]byte, 0, sizeMsg)
	// build a string component by component
	return string(b)
}

/*func main() {
	fmt.Println(msgPart1)
	stringNumber := createMsg4wat("Netmagic", "176500")
	fmt.Println(stringNumber)

}*/

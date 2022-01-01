package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := input
	var ids []string
	scanner := bufio.NewScanner(bytes.NewBufferString(in))
	for scanner.Scan() {
		l := strings.Trim(scanner.Text(), " \r\n\t")
		if l == "" {
			continue
		}
		ids = append(ids, l)
	}

	var dup2 int
	var dup3 int
	for _, id := range ids {
		dups := countDuplicates(id)
		//fmt.Printf("%q => %v\n", id, dups)
		if _, ok := dups[2]; ok {
			dup2++
		}
		if _, ok := dups[3]; ok {
			dup3++
		}
	}
	fmt.Printf("product: %d\n", dup2*dup3)

	for _, s1 := range ids {
		for _, s2 := range ids {
			diff, common := diffChars(s1, s2)
			if diff == 1 {
				fmt.Printf("found pair: %q, %q => common = %q\n", s1, s2, common)
				os.Exit(0)
			}
		}
	}
	fmt.Printf("found no pair\n")
}

func countDuplicates(s string) map[int]int {
	cnts := map[rune]int{}
	for _, r := range s {
		cnts[r]++
	}
	dups := map[int]int{}
	for _, cnt := range cnts {
		if cnt > 1 {
			dups[cnt]++
		}
	}
	return dups
}

func diffChars(s1, s2 string) (int, string) {
	if len(s1) != len(s2) {
		return -1, ""
	}
	var d int
	var common string
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			d++
		} else {
			common += string(s1[i])
		}
	}
	return d, common
}

var inputTest = `
abcdef
bababc
abbcde
abcccd
aabcdd
abcdee
ababab
`

var inputTest1 = `
abcde
fghij
klmno
pqrst
fguij
axcye
wvxyz
`

var input = `
luojygedpvsthptkxiwnaorzmq
lucjqgedppsbhftkxiwnaorlmq
lucjmgefpvsbhftkxiwnaorziq
lucjvgedpvsbxftkxiwpaorzmq
lrcjygedjvmbhftkxiwnaorzmq
lucjygedpvsbhftkxiwnootzmu
eucjygedpvsbhftbxiwnaorzfq
lulnygedpvsbhftkxrwnaorzmq
lucsygedpvsohftkxqwnaorzmq
lucjyaedpvsnhftkxiwnaorzyq
lunjygedpvsohftkxiwnaorzmb
lucjxgedpvsbhrtkxiwnamrzmq
lucjygevpvsbhftkxcwnaorzma
lucjbgedpvsbhftrxiwnaoazmq
llcjygkdpvhbhftkxiwnaorzmq
lmcjygxdpvsbhftkxswnaorzmq
lucpygedpvsbhftkxiwraorzmc
lucjbgrdpvsblftkxiwnaorzmq
lucjfgedpvsbhftkxiwnaurzmv
lucjygenpvsbhytkxiwnaorgmq
luqjyredsvsbhftkxiwnaorzmq
lucjygedpvavhftkxiwnaorumq
gucjygedpvsbhkxkxiwnaorzmq
lucjygedpvsbhftkxlwnaordcq
lucjygedpvibhfqkxiwnaorzmm
lucjegedpvsbaftkxewnaorzmq
kucjygeqpvsbhfokxiwnaorzmq
lugjygedwvsbhftkxiwnatrzmq
lucjygedqvsbhftdxiwnayrzmq
lucjygekpvsbuftkxiwnaqrzmq
lucjygedpvsbhfbkxiwnaoozdq
lscjygedpvzchftkxiwnaorzmq
luckygedpvsbxftkxiwnaorvmq
luyjygedgvsbhptkxiwnaorzmq
lmcjygedpvsbhfckxiwnaodzmq
lucmygedwvybhftkxiwnaorzmq
lgcjhgedavsbhftkxiwnaorzmq
lucjugedpvsbhftkxiwmaoozmq
lucjygedpvybhftkxkwnaorumq
lucjygedpvzbhfakxiwnaorzpq
lucjygedpvsbhftyxzwnajrzmq
lucjygedpvsdhfakxiwnoorzmq
luyjygeopvhbhftkxiwnaorzmq
lucjygadpvsbhntkxiwnaorzmx
lucjygedzvsbhftkiiwuaorzmq
sucjygodpvsbhftkxiwuaorzmq
euijygydpvsbhftkxiwnaorzmq
lucjlgeduvsbhftkxicnaorzmq
lucjdgedpvsbhfgkxiwnhorzmq
lucjymedpvsbhotkxiqnaorzmq
lucjygmdpvsbhftkxywnairzmq
lucjggedpvsbhfxkxiqnaorzmq
sucjygedpvsbhftkxiwnaorjmv
lucjlgedpvsbhftkxiwnairzmg
lucjygedppubhftkxijnaorzmq
lucjyxedpvsvhftkxlwnaorzmq
lucjygedpvxbhftkfiwyaorzmq
lucjygedposbhftkniwnaorzmw
lucjygewpvsbhftgxiwnavrzmq
lucjynedpvsbmftkaiwnaorzmq
lucjyhedpvzbhftkxiwncorzmq
lucjygedpvsbhfikpiwnaoezmq
lupjypedpvsbhftkjiwnaorzmq
lucjygudpvsbhfwkxivnaorzmq
lucjygrdpvsbhatkxzwnaorzmq
lucjbgmdpvsbhftkxihnaorzmq
lucjmgedpvpbhftkxiwnaorcmq
lucjygedpvskhfukmiwnaorzmq
lucjygedgvsbhftkxiwnvprzmq
lucjzgedppsbhytkxiwnaorzmq
lfcjypedpvsbhftrxiwnaorzmq
lucjyqldphsbhftkxiwnaorzmq
lucjygedpvsbhftzxewnaorzqq
lucjygeapvsbhftkxiinoorzmq
lucjygedpvszhftguiwnaorzmq
luojygedpvsbhftkxawnaornmq
lucjygedpcsboetkxiwnaorzmq
lufjygedpvfbhftaxiwnaorzmq
luciygedpvsbhftkxhwaaorzmq
lucjygedpvnbhftkaiwnaorzmc
lucjygedpvsbhftkxiwcaorbdq
lucjygelpvsbhftaxiwsaorzmq
lujjygedpssbhftkxiwnaorzmr
ludjygedpvsbhftkxiynaorzmj
lukjygeedvsbhftkxiwnaorzmq
lucjqpedpvsbhftkxiwnaozzmq
jucjygedpvsbhftkxgwnaorqmq
llwjygedpvsbhetkxiwnaorzmq
rucjygedpvsbhftkxiwndorymq
lucjygedpvsbhftvxswnaorwmq
lucjygerpvsbhfykxiwnaormmq
lucjynedpvsbhftkxijnaorziq
ljcjygedpvrbhftkeiwnaorzmq
lucjygedpnsbhftkxiwhaornmq
lucjygadpvsbhftkxibnaorzqq
lucjqgedpvsihftkxiwnaorzdq
lucjygedpvsqhfttjiwnaorzmq
llcjygedsvsbhftkxiwwaorzmq
lfckygedpvsbhftkxiunaorzmq
lucjyeedpdsbhftkxiwnaotzmq
lucjygedpvsbhftkoiwnaoqzcq
huwjvgedpvsbhftkxiwnaorzmq
lucjygldpvsbdhtkxiwnaorzmq
lycxygedpvsbhftmxiwnaorzmq
lucjygedpvsbhftyxianvorzmq
lucuygedpdsbhqtkxiwnaorzmq
lucjyggdpvsbhftkxiwnavremq
lucjyggdpvsbkftkxiwnaorbmq
luchyqedpvsbhftixiwnaorzmq
lpcnygedpvsbhftkxzwnaorzmq
lucjygedpvsihftkxiwfaortmq
lucjygvdpvsbhgtkxiwnamrzmq
lucjygodpvrbhqtkxiwnaorzmq
lucjygedpfsbhftkxipnaorzma
lucjygedpvsbhftkxpcjaorzmq
lucjygodbmsbhftkxiwnaorzmq
lucjygedpvsbhftkxipnaogzmb
luxjygjdpvsbhltkxiwnaorzmq
lucxygedpvsbhftkxzwnaorjmq
luajygedpvsbhftzxiwaaorzmq
lhcjygedpvsqhftfxiwnaorzmq
lucjygecphsbhftkxiwnaprzmq
lucjygedpvsbhptkxifnaorqmq
lucjygedpvichftkpiwnaorzmq
lucjygedpcsbhstkxswnaorzmq
kucjygedpvsbhftkxiwbyorzmq
lfpjxgedpvsbhftkxiwnaorzmq
lucjytldpvsbhftkxiwdaorzmq
lufjygedpvfbhftbxiwnaorzmq
lucjygebpvgbhftkxipnaorzmq
luujygedpvdbhftkxiwnaorzmd
lucjygedpvsbhfbyxwwnaorzmq
lucjygedpvsbhftkxiwnaoqpmw
qucgygedpvsbhftkxiwnaortmq
ludjtgedpvsbhftkxiunaorzmq
lucjyiedovsbhftkxiwjaorzmq
lucjygedpysbjftoxiwnaorzmq
lumjygedpvsbuftkxiknaorzmq
lucjygedpvsbhfokxgonaorzmq
lucjygeqpvsbhftkfiwnaorzeq
lucjygedpvskhftkxiwntorkmq
luujygedpvsbhftkxiwraorzmt
lucwygedpvsbjftkxiwnaorzmj
jucjyfedcvsbhftkxiwnaorzmq
luujygedpnsehftkxiwnaorzmq
lucjygedpvszhfckxiwnaorzmi
lucjyredpvsbzftkpiwnaorzmq
lucjygedpvsbwfgkxiwnaorzoq
lucjygedpvgbhftkpiwnaorzms
lucjygedpvjbhftkxzwnaoizmq
vucjycedpvsbhftkxiwfaorzmq
luawygeapvsbhftkxiwnaorzmq
lucjygetpvsbhftkxiwnaafzmq
lucjvgedpvsbhftkxywnavrzmq
luolygedpvsbgftkxiwnaorzmq
likjygedpvsbhftkxiwnabrzmq
lucjygedovsbhftkxirpaorzmq
lucjygedphsshftkxqwnaorzmq
uuqjygewpvsbhftkxiwnaorzmq
lucjygedcvsbhftkxiwoarrzmq
lucnygedpvsbhfakxiwnaorzms
lucjygedpvsbhntkxiwnawrzmb
lucjygedpvsblfxkxivnaorzmq
lucjygedpvsghftkxiwnaawzmq
yucjygedpgsbhftkxiwnaorzbq
lucjyweapvsbhftkxiwnaoezmq
lucjygevpvsbyftcxiwnaorzmq
luejygedovsbhftkxiwnqorzmq
lucjyqedpvsbhfbkxiwnaorzms
lucjypedpvsbhftwxiwnhorzmq
lucjygedpvsbhmtkviwxaorzmq
lucjogedpvpbhftkxiwnaorqmq
lucjygedpvsbhztkxkwnaoazmq
lucjyaedpvsbcftkxiwnaorzhq
lucjygbdpvkbhftkxiznaorzmq
lucpygedpvzbhftkxfwnaorzmq
lucjmgedpcsbhftkxiwnaoezmq
lucjygedyvsbbftkxiwnnorzmq
lucjyyedpvsbhftuxiwnaonzmq
lucjygfdpvsbhutkxiwnaorzmt
uccjygedpvschftkxiwnaorzmq
lusjygedpvbbhqtkxiwnaorzmq
ducuygedpvsbhftkxiwnaorzyq
lucjygkdvwsbhftkxiwnaorzmq
cucjyyedpvsbhftkxiwnaerzmq
lucjygedavsbhftkxiwnkorzbq
lucjygedmvsyhftkxiwiaorzmq
lucjygeipvsbhfpkxiwnaorzpq
vucjugedvvsbhftkxiwnaorzmq
lucjyzedpvsbhftkxpwnaoozmq
lucjygedpvgbhftkxiwtaorzqq
lecjygedpvcwhftkxiwnaorzmq
lucjyghdpvsbhfcyxiwnaorzmq
lucjygedpvesqftkxiwnaorzmq
lucjyjehpvsbhftbxiwnaorzmq
lucjygedpvtbhdtkxignaorzmq
lucjygxdpgsbhftkxivnaorzmq
lucjygvdpvsbhftkpiwnaorzqq
lucjysedpvsbhftkxiwnalrzmc
lucjygedpvkbhjtkxiwnaorsmq
lucjygedpvsbvfgkxiwnaerzmq
lucjygedpvsihftkxilnaorzmu
lvcvygndpvsbhftkxiwnaorzmq
lucjysedpqsbhftkxiwnaordmq
lucsygeypvsbhftkwiwnaorzmq
lucjygewpotbhftkxiwnaorzmq
lucjysedpvsbhftkxiwnanrzmv
lucjygedpvsbhutkxiwnaoplmq
wucjygedpvsqbftkxiwnaorzmq
lacjygeepvsbhftkxiwnjorzmq
lucjygedpusyhftkxicnaorzmq
qucjyredpvsbhftkxiwnworzmq
lucjygedevsbhftkgiwnayrzmq
lucjygedpksbrftkliwnaorzmq
lucjygedpvsbhfgkxisnaorzeq
lucjygedpvhdhftkeiwnaorzmq
lucjsgedpvsboftkxiwnaorumq
luctygedpvsbhftouiwnaorzmq
lucjygedpvsjhfukjiwnaorzmq
lucjagrepvsbhftkxiwnaorzmq
lucjkgerpvsbhftkxiwnairzmq
turjygedpvsbnftkxiwnaorzmq
lbcjygedpvsbhftkdpwnaorzmq
lucpygedpvsbhftkxnwnoorzmq
jucjygedpvsbhbtkxicnaorzmq
lecjygedpvsbhftkriwnaogzmq
licjyvcdpvsbhftkxiwnaorzmq
lrcjygewpnsbhftkxiwnaorzmq
ltcxygedpvlbhftkxiwnaorzmq
luctygedpvhbhztkxiwnaorzmq
lucwygedplsbhfakxiwnaorzmq
lucjygedpnsbhftkxiwjaoezmq
lucpygedptsbhftkxiwnaorzmo
lucjygedpvibhqtkxiknaorzmq
lucjwgqdpvrbhftkxiwnaorzmq
lucjmgkdpvsbhftkxiwraorzmq
lucjygwupvsbhftkxiznaorzmq
lucjhgedpvobhftkxiwncorzmq
lucjygedpvsbhftkxiwnaohtmj
lucjygedpvsbeftkfiwnaorzyq
lucjygcdpvsbpftkhiwnaorzmq
lucjygedpmsbhftkxiwnkouzmq
oucjygedpvsbyftkximnaorzmq
lucjcgedpvsbhftkxywnforzmq
lfcjygedfvsbdftkxiwnaorzmq
ducjygedevsbhfttxiwnaorzmq
ldcjdgedpvsbhftkxiwnavrzmq
lucjymedmvsbhqtkxiwnaorzmq
lucjygedpvabhftkxiwnasrlmq
lucjygefpvsbhftkxmwnaorkmq
`

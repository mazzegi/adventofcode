/*
--- Day 6: Custom Customs ---
As your flight approaches the regional airport where you'll switch to a much larger plane, customs declaration forms are distributed to the passengers.

The form asks a series of 26 yes-or-no questions marked a through z. All you need to do is identify the questions for which anyone in your group answers "yes". Since your group is just you, this doesn't take very long.

However, the person sitting next to you seems to be experiencing a language barrier and asks if you can help. For each of the people in their group, you write down the questions for which they answer "yes", one per line. For example:

abcx
abcy
abcz
In this group, there are 6 questions to which anyone answered "yes": a, b, c, x, y, and z. (Duplicate answers to the same question don't count extra; each question counts at most once.)

Another group asks for your help, then another, and eventually you've collected answers from every group on the plane (your puzzle input). Each group's answers are separated by a blank line, and within each group, each person's answers are on a single line. For example:

abc

a
b
c

ab
ac

a
a
a
a

b
This list represents answers from five groups:

The first group contains one person who answered "yes" to 3 questions: a, b, and c.
The second group contains three people; combined, they answered "yes" to 3 questions: a, b, and c.
The third group contains two people; combined, they answered "yes" to 3 questions: a, b, and c.
The fourth group contains four people; combined, they answered "yes" to only 1 question, a.
The last group contains one person who answered "yes" to only 1 question, b.
In this example, the sum of these counts is 3 + 3 + 3 + 1 + 1 = 11.

For each group, count the number of questions to which anyone answered "yes". What is the sum of those counts?
*/

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

func main() {
	buf := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(buf)
	curr := []string{}

	var gs []Group
	flush := func() {
		if len(curr) == 0 {
			return
		}
		gs = append(gs, Group{
			Answers: curr,
		})
		curr = []string{}
	}

	for scanner.Scan() {
		l := strings.Trim(scanner.Text(), " \r\n\t")
		if l == "" {
			flush()
			continue
		}
		curr = append(curr, l)
	}
	flush()

	var sum int
	for _, g := range gs {
		sum += g.UniqueChars()
	}
	fmt.Printf("unique-chars sum: %d\n", sum)

	sum = 0
	for _, g := range gs {
		sum += g.CommonChars()
	}
	fmt.Printf("common-chars sum: %d\n", sum)
}

type Group struct {
	Answers []string
}

func (g Group) UniqueChars() int {
	m := map[rune]int{}
	for _, a := range g.Answers {
		for _, r := range a {
			m[r]++
		}
	}
	return len(m)
}

func (g Group) CommonChars() int {
	isContainedBy := func(r rune, a string) bool {
		for _, ar := range a {
			if ar == r {
				return true
			}
		}
		return false
	}

	isContainedByEach := func(r rune) bool {
		for _, a := range g.Answers {
			if !isContainedBy(r, a) {
				return false
			}
		}
		return true
	}

	var cnt int
	for _, r := range g.Answers[0] {
		if isContainedByEach(r) {
			cnt++
		}
	}
	return cnt
}

var inputTest = `
abc

a
b
c

ab
ac

a
a
a
a

b
`

var input = `
adgvrhblps
pghsdrbmalv
hrlbpdasgv
bgvsdplahr

lgnpfhrm
hwmng
gunhmo

txkeafsbgjuizd
etmcgdbfajuz

xdtzjioqavmchsbfrkp
bzjkriqmvxedotpcf
azjckxmqovtidbprf
fcxmevrkojzpdibqt
roifztvxmbpwcndkjq

ylcixrdoejn
dyoejlrcxin
dnorleijxyc

u
u

nfjrt
hvgqxculeszok
pmwryfdiab

gh
ebzkr
byqusikr

b
b
alyr
t

jzbafspmynwgqdeuklxc
anbkgduyexjlzmpcfwq
zxdkiujqeglcfynwpamb
qwcubamjpxyngkdlzfe
upadgfkywmnxeqbhjzcl

g
l
g
g
pg

mazi
mnzdai
zmia
izma

pqzl
pzql
zvciplq
szpqrl
plczqi

mquiosghkwtjz
tgnewskmjuozqi
gjwzmqcntosuik
ybtjoqzuwmfgisk

nvwocedxyqiahgt
zbjvipyfeacuwrt

ojkhxmwqs
kqxhmrzjwo
mqkjxhwo

t
t
t
t

dhrxaqejfng
fhdlbxejqra
fqaeldhjrx
seaqjxzdfthr
adjhrqlxef

t
bczrx
ueit
nu

idtmkhnlpbsqfxcevjry
pricynubfvdhjqesmktlx

wt
nw
w

mhqe
ey
e

kwclmuovtihabd
aqheucldfwno
adwnhoclqugez

lbwnohgs
osnbghw
swogbhfn
hnbosgw
gsowhnb

jsxyihpaotdlvckgeznqfbmwr
xkhfnwitojasqyrgcelpmbvd
synvibmdpfxwqltgcjohkare
grxicqelsovtdfhujwkanypmb

druazhicbm
spqizhucrdabmv
imucbayhgzdr

fnx
unqx

zjm
qfwrycva

uktvygcmf
yfktcxmv
ctlkvqysfmb

j
jng
j
j

t
tw
t
t

vogs
ogsv

tzgqxfauw
wgfxqkuaz
dkzawuftgq
gbuwaqzfjo

iphexunlbs
tayqgzmkjfdcrvow

fcdplxkhjbngwzqmiry
khuyrlgvxnpfbdmosqizjc
djzyxikqnmbrlgfhcp
jxycgeidnhlkzqpbmrf
xckrhgbzfnjlpdmiyq

lezco
ezloc
zolec
ocelz
oeczl

emgluf
flebmhg

knafibd
kbrnfd

xktmidfrz
fajlidmrtzy
wftzdm
omeqbfgdzpnut

tndpxaveifhcmzyjrsqlub
oyanuslpqdcevfih

hblawudzknspf
gxnstrmiyqpejo

lesyaikhnvoguwftzrmqdcpx
pxuqoayiwjzvglstrhfkmcend
uvyfepmxokdrhzlcsniagtwq
ynfmkwxqalpdrezhoguvstic

grscfn
azlgdcn
ovcgzn
lwxgctan
mgkbiyjhpnc

ji
ji
iaj

ivtkp
ikwp
ipko
pkiw

p
p

yushxnqa
uokzcaxni
xalvuyn

elw
wle
wlre
wleq
uwel

bvfxe
ebvfx

roaygfv
udazcljhy
pbaryv
aptxy

zdpkslmvnhyfua
czhfkylumrvgoqaspb
twenpsyzkjmihavuflx

zsv
bnzjfcamkw
rvdzq
ltypirz

q
w
q
q
q

uic
uic
uic
iufc

rdve
cvzre
fqls

gxmnabhd
xdbwjngmra
gmxebadny
bgahdxomn

gnoyc
olgy
poyfjg
wsdouiyghqe
jyog

q
e

adszvgclfomph
qavsicdlghozpfm

czfhs
hzfo
zhf

luzrpeivwagyfktjhnob
tlwynevkojgpbfuihr
gklbpjvyoituefrwhn
vhypwujfrbktonedigl
byknejlowfrptghivu

xbu
xbu
xub
bux
xub

vcqtxyf
tqcvfy

vumkixjhaozgfbneydpr
yxeazjmvbnhukpigodfr
pjnomhaxuvdkbiryzfge
oezrxngvpajfkdmuyibh

mkgdruoiawfh
dsumegkpjoq

oucrtxbqienfvsmlzjkdwgpa
akrqisvxbudelwgpmnozcj
xreauzqglpcvmoidbskwjn
mwgpljcusakrbozxvndeqyi
jdrlvcimoewkpxgqasnuzb

dshfqywgbc
eiqflapjh
vrmjfhnq

kojresxi
ubclnpghf

kuitbfh
ofdkqt
fztk

lsawunycojgvi
ornsyijx
yojpdsnbxri
tjysrinoe

ezuqslamgi
alimgeuqz
amgfqzluiew
nmalugivwqze

lrm
mr
vmrgt
mr
mlr

ybcnizhwt
ycbmztih
octgiqsyhxz
hktzcfyilew

mqigxepzwudokytanf
ayztcxbigunordqlmhk

wnfy
wzfp
fjlewksb
pcfwyn

o
c
na
qku

zg
gtl

syr
pbnmqc
j
sk

abenmqvsyotzrcjwglfdx
cnmqrxatjvowbylfgdezs

kfzlgphuwqv
levuhkwqcpfg

iazxrty
xryz

ikxcpebaruwhlm
iqtxadbeyckuvsh

wkfsrvgbyjxpq
rqbyxwfpkjgsv
fgwsyjqkxovrbp
frpyqxbjvwsgk

po
o

arbnupkvxqodcg
znpobhraxugvdckqt
mkaupgobdvscnqrjx
cbqknjavpgwxuord

ftpka
vtifcl
dfwusmbyrhto

y
q

drsleaqghypbtinmfuo
gzinmrwudjbpyloefqahs
fiomngyurlpheqdabs

v
bx
c
v

kh
zk

qkzicvjbgawfdpxmtsloye
ylofxwpvkqemsaigzjtcdb

qf
fq
raf

cbvxjytmsa
tbmxcvyjsa
ybmstxjacv

aekpuyliz
eqpyxls
rwyepchjflomg
oedybpl
evnyclxqp

yxtfnavdhopmcw
qlugbvjcz
xvcnymkrh

yksamve
yksmvxea
mskavye
yamdswevk
eksayvmx

p
k

psrgzxmaif
ifraxgzo

aehoplsvfqr
jfqvkpoas
vsjfxpqcoma
asmfoqpv
uvospqaxf

ercgdyism
msrvebzc
mgrescq
desfrmico

dborguxznkyqciavtfs
kzoymfndrsaihqtcbvgxu
anbcwfrdgjxktlyivouzsq
bgifuzsrekocxydqtnva
dtpsikyavoucxqgbhfznr

jvyhfbpxtkizmlesw
fvmjhkesbpzywixtl
btlezmshjwpxiyfvk

jawdspkghblyvornxmicuq
nrujdhxokcwgylimvqpba
lcrnhktgvipyubwqdomjax

vhiatyrwksodemcxn
hdsamtvuikjqecnwryox

ixthorcaevkp
xiyhrknlboczwjfm

ck
kc
czkmyet
kc

wxjqkpgnshue
qojwkefphzcrgaxn
ektphqjxgnwu

raksq
mdu
a
asb

putzomyciwaqkvbfsldng
absdvyntgikruqpzoflcmw
dkscznaiwulvogqfybmjpt

sxbqjv
jqvxbs
vlqsxjbo
xsbqjv
jvqxbsy

ynbakgvc
yvtnlbsjg
vrqxephwfby

xhdjkrpltfoe
qjplehkotdbfrx
kxfdaotsperlhj

ewgsmi
jafkti
yqkf
copzldurbx

xc
ex
r
rtwce
slk

zimewg
nmzrj

qylbm
byqml
lqmyb
mxyqlbr

egsktr
erskdgt
rkesgt

zkgqireysupwjthdb
swkoyzdcqibpvnr
kyvmwzpirdqsfbclx

jvytcgmkeqalzonhx
kvcjnwztxdmlyoaqg
lznqaxcvjytogkm
vabgfulyxkjoimqztcn

q
q
q
fh
q

xykitfhn
xnhytzilf
tiznfhx
tfxihnre
xhilzytfn

zmxyoakgt
yawiomrgsd
oafgmyqr

lpgqshiyuxjvabz
lzbuhyiqsjagepx
urisxlzjabpyqgh
xsuiqljypzgabh
gkylpamihsotqfdxzcbjwu

jhf
htdejm
jtabh
cjeigmzh
suljqyorxh

hcjkbvmlaoen
faojknvxmehblyc
nhmovlejcbka
vonhckilembja
kojnmvebtclha

clamzrewgonksuqvpb
zusorlkmqpnvwaecgb

uka
ea
aw
a
lxa

lg
lwtg
tgl
glw
rlg

vfsxidctgah
dvhcftsxgia
vdhxgimfsant

tschpyqa

ckn
xckznb
mnsvlfqaitjpdk
cwkzn
kzoyenb

ro
ro
or
or
oer

icumvltoqafdy
tfsxarpnh
atsfgw

btfenmovgusqzijxpcrhdakl
diapqfxrnevlzutsycjgbkom

utvbhgw
vbhgtwu

nkfgajmvbpyowthxr
cwxrtknpmgfazjoy
rwxmfatgouckjypnd

kdlubgiap
podcivgyueafkhm
purgnkixazd
sukdgixap
udkpglia

ajnumbsgr
ibmjsrkagc
bsgozmra

vkbso
ykglhto
gmodikp

sxupqv
xvuqps

x
r
c
x

fjoyuhztseqrawpciglx
igfeplsxtrhaqyzojwuc
rpwgucisoyhflaxtzejq
yqzjphegfuxltawicsor
wzahifloypuxegjscqtr

xeironzdugvqc
quevgsfronxzc

wdjrsknlbihcqzxuv
bzlxjnukvhrsdicwq
cuqiwsxkzvjbrndhl
ubclivhszjdwqrnxk
jldqkihxrzuvsnbwc

dzklrogfvmn
xuwpaqcetij

d
o
m
d

xhgzbj
jvbhg

snbcamefqhuovxkjpz
htckmpybdgvrsaoeuj
pkcbewsgvojmhauidl

i
k
tqj

koyawh
auo
qaoer
zrapo
ioa

broqhaylgsmxzkjtuvf
hatgluorzvkyxsmfjb
whytbkrvsumgzxoalfij
tzabxgfrqnusomvhjykl

hoplnck
ayqxdjofr

dmjspac
dhoazjp
idxfrwepyjo
qlpjgtvndub

riabvstyxphgoez
projxgyhietsavz

jhfqikpmz
kfpiyjvmzsqh
vphkqizjmsf
wpqfmlzkjheid

diap
zkhacy
waog
fojgas
roa

rujzxoygliqnmvkpc
vlonrgpjqzyxmfkic
zgxqkjocirpvlynm

lzgja
jlza
kjazrcl

vteiqrpkjybsmndhfl
lsndjpqertkvhbimyf
lpjmnbyqshvtdfirek
vpibqrehmsnflykjdt

iwqlhotpjsa
slhjyuiotrdk
szjhlocimt
tiajhwsofl

poz
dp

yeqdmto
zgtqyvoli

ahkfsocmdgnwp
cpwhjsmgodakf
clhudwakftvogsepi
dkargcowsphf

p
p
p

eqbcr
jceq
xeocqmi
fqwcve
mjebocq

xkei
ivuaexqk
kxie
ikxe
kxei

jsmzk
cmdpyf
xshm

hdwlcgzyuvke
gjlaqpry
rlgomxy

comszwhfrqlpnvytei
iqpvwolsmnfhray
oulphgswymvfrqnix

lcisfxnzwy
csixywfbnlz
zelcfyiwsxn

lmrsia
gowals

lyknfogpu
zoiw
tesjova
moes
odi

ktfajxloz
ojaxzflk
akzlfoxj
fxkolazj

iwxzmukbhoy
xlzkymbiuwo
bkuzomwxtiqyd
wmokbizxyul

kr
kr
skr
kr
rk

bdpfgnwxtk
gwtkdflnjpzx
kngxpfwdt

xkhmzvylowsgucrfepjai
xriqvhjlozpdkuemcfa

avxoebdgzmqphfytcn
cvfrmneolxtyzdgpqh

wjvslhmufgqxcnodz
fswclgdvxnumo

t
uqi

xopmekybt
pkrxytebm
mtpxbyek

vkxcrizwuhtad
wairuvhdmzcxkt

vrlpifawedksqzhcgnybtjxo
qoplwufekztbhdnvxgarjsciy
svzmteygfbkrlxohjqcpaidwn
rakocwstbpidlhxjgeqzfnuvy

o
o
o
o

ihzarybenqmtpg
zqbiyhanpgetrm
hqizynrpbagetm
yhbgrzeqnipatm
gteahmprbyqinz

jv
uxkasvt
jrgv
qv
ivbpld

sxigewcouvbj
xjwobcsugive
cbjvxsuowneig
uxobwigsjcev
ejogiwbxvcus

enmvubp
pvubm

qrjythvlzubgeikafo
necadbmrtzulvwypx

rbuslgk
lgsbuzkr
lbkugrs
sgkbulr
rublgks

whb
wbh
whb
bwh
whb

khtcw
wkcth
wtkch
ckwth
kwhtc

mkzdbo
dkjz
kbnzlmd
duytksvzgq

brkxvtjpdhqw
ntjqizashkpvwd
dqfrptwkshmvzj
phegdjolwvkqtu
dkhnwqacjpmtvf

bdrioxsqahp
adishrbqoxp
oiqpsbdtyrha
dapobrxsqih

oqsgiptmzlwcxkhvrfedu
rpwhtlgkcuidvqfosazxjem
rmgkeisxzhqwopltcfvud
vfcxiruwlhgoszmekqdpt
hmprfdvoxetulgizsqwkc

luzam
mwuplz
mlurjz

xlpmhzudwriytscobnkfq
uinamwhrtpobfvqceyd
ibflutdsohnqrpcymw
nqprtwuyfdcmiohb
ufhwmrinodkcbtpjyq

ns
ns
sn
sn

dufzbegpjrsaiqcthv
pfzqdjhatbuvescrkig
tbehcpfsvagdqirjuk
tebiravqjpfushwgcdx

gpjafzel
mzgftlead
bleynhgvfcs
egdqfalmz

zbarcdwiukgoqhmtnyjps
sujqytgrbnwzdmhalicok

fyvgdcbqx
fqcgdyin
etfmoqgdsjyc
gqfcdy
fyqcngd

am
bjm
sdm

frvkaupe
qzipevkya
vzabpkesd
bevpymka
pkxeva

t
t

dcveyjgtfkbuo
ebogjfydkvcut
gkoyfeudtcvjb
mvygqckjftuodbe
tujbkeocdgvyf

bkthnlgp
tnghb
gthnb
gnhbt

hpkybmeowacsrgnijfl
jzrsgyuwickpfbmolneh
hcflvrgqpwmdoknsjeyib

qxincam
zdjrlksqhgt
cbafqwo
vqyei

izdkxrpuqgjte
asyngkxrudpvmie
bdkwpgmexruci
opdiknuxger
gecixdrpluok

jymuwvfrixclhga
invyhuglxmcjw
imwghculxvjny
gjmlhiwvyctux
yvjhixugwolnmc

xbcwhv
chdwgxbv

ujycfnsvgqkpwheal
erugwaykpqcnhomlf

gyk
glhrk
gwk
qowygk

vcnju
vuoijbczt
dvcujg

rlsy
ryls
ylsr

tjn
t
n
rixz
t

i
i
i

elq
qle

ksahwvdjbtfle
lfksaebjhtwdv

qfacst
qfaskg
fqasoktb
vmfuilqzprasd

apuxh
uxpah
onahmwupxb
hxpqau
huqapsx

hjbnwua
huwbranj

dj
z
h
hz

hzqo
oqz
uwsozq
zyqxo

ckotpfx
lhrcnpmgfviwobd

wecahmszdjnoix
hmxaijswzydtcgn
njcqtmzixsadlh
cudmxzeasjnh
mxzahjckbnpdvrs

lsqahvixforkw
fxiqskaohw
afjhnxqowsik
xstiqkhmoafpbw

htdcbm
ex
sxq

pfgkrjhub
nptjfixgraubykhld
pmufkbrhjvg
jhrfbgpvuk

sptnugvdfmckehw
wknfhmsucertpvg

zolicbnrvtdwhy
oqdunzywp
zodykuwxgspn
wsanojdzxqy

druisbhpckvwofna
wvcboafsnhipukd
bpsadlcvnkhfziuo
bqfujsvmhcoakitdpne
visaodpfuchkngb

ewniqoulvgstzchkxdmpyajrbf
cbyihmnuradvlktxpjwsogzfeq

zgsdqwj
qdzsg

ombl
yeulqbnimw
hljodbmk

wtzkorqhupxby
yzxhtkwpord
skwzpjtgyrxho

mstgrlwo
rpyowgdts

pidn
ipesrfv
owpdim
pibt

tu
tu

pcdatkohjgbelmqvfxwnu
ukevynsprjtclaxhgmdqfwb
twqungiobxflhckapedjvm
pgqvtowflujnmadbeckxh
damfntwkjlcvhbqepxgu

ilksfovqamrhngcdjyexbt
htkvnbrqaefsloxgciymjd
qfgvrxmbaoeldnkyiswjhtc
ftxpoyhbjkcevdasmigrlqn

oq
qdo
sdlqoa
eopbizq
sndqo

iblpgzoekrnmwhjstay
tsropbkwgyaexilujn
salrwjgbtikpeyon
bsyirktpjonxualewg
teaqljpowyrigbknsd

d
d
d
d

zgqmhnlvekcyutxsbifw
zmkqgwvyjrcdfbu
ckoawvyzgumrfbqd
fdqbypumcgwvzk

krqlncwuxvgbdpsmyjfa
lmuybxdqsvkrjfcpwnga
zkbqnhjspdcumrxgaelyfwv

vafsedqzyrtnjpwug
gvcustypjedaw
owpydujgtaivsmke
tvadlgweiypsju
jebvtoysphugadw

pyixcohfrvtenaugsdmjq
catjqrivdfseogunymx
tvquyscedoxrgifjman
qagtfiocrjmuvdsenyx
gmeufoqaxcrnjdvstyi

p
p
p

jagycukemiqdxw
cmuyadqgkewixjl
fqawimgzjkyuxbdec

mahdekwibs
admswkhei
debamswk
maovdyeskw
mbsdkweaf

wfqimrc
rtqhcwlgdxnkvpo
arwysquc

kroumefnixa
lcn

hvaugnsf
nfxqkugas

wheyqtmdvljucnsbia
hvmyteblnuqacw
ehnultwbyacmqv
mlybuhqtcewvna

pxhafcvymglkzoq
xhwkctlgabz

or
jodr
or
gro
forha

fzmohqld
szqf
iqawckfrp

hkxuwmsbdrpt
ptysjhxkwbmru
tpmbsrkgwh
kphmtrusvwbq
krmhpftsbwzn

r
r
otr
r
r

sxtlkvjacpbngoziewq
jiskzqanlowbxpvtge

r
r
r
r
r

bqwsz
wsfb

zwvyqaichxjre
whvayjrzicxqe
arxcivhewqyzj

mpzasbnregjdwlxuqoft
dnxlrwubzkaosypq

gfaizbqlchy
zibatqykfu
bsufzjqxyiam

fsy
sf
efs
fs
fs

asmugcxhbekl
dkhcluwsmegaxb
kubhcmalegsx
cgluxmbhekas
bajrhelscgumknx

qdpwfec
wjdqfecy
cwdefqj
wqvdclsfe
dcqpfew

r
jcowmyplv
itdru
shbfr

soublfhe
lksfzjdymaw
ywfnrskcl

zsybelirjfwm
sfxiznhmablgtj
jbcifusqlopmdz
znhjerlkfmybis

knvmejaydo
mjoaevkw
jkmhcoave
agomevkyj
elqftajmoskvi

ziajfk
zakifj
kazfij
zfajki
fziajk

ujzyfat
tyuafjz

xihdrjyzvl
axujhtldyr
xdohklyjr
xrlejhyd
bxolyrfkjdh

i
i
i
i

ocnxzlprutqmhsij
xgmloizutqj
xmtlgioujqz
guoxlamzitjq

jhoepxcuwvn
xeuvnfaodjwpc
upjoxwcevn
xncelrvojupwg
qeusjvwynkpotxc

etdilxy
lytdiex
hdeltgryiw
lyixvedt

rge
gr
grp
gr
grp

iquzsnh
qsfunwh
qtsgprvhcn

lipmbftsh
myho
uqzngarkvw

ybrtackvdgqsmonlezuw
ywlzmsgbtocakedrqnv
rbceqauytzgmowkldvns
tyncgkldqserbjoavwmzh
oydsvatnmwzgceqkrlbx

qjb
irgfqjope
jkq
jqb

ncjlbpzedywkshfmqvagxi
jfipaqbmevkgdxnschztl
gbvxnclahmsiqjdekzfp
kedsxbzcvijglphmqfna
szcadkbnpeghjvqifxml

nfteow
ytewforn
estonfuw
erwtnfo

vhawcorfpxydzieqmbg
opyqhzdabfwcegrximv
vbpfqzcierhdgawoxmy
dvprwqzfioagxbycehm
yhdvpxzimqeorcabgwf

bfnrkixsut
jhywbrncziaxkp
iodknxrb
xbrknli
xkfironbu

xwpqbhyakf
rxayi
ogaxy

tnwl
wbnl
lhwn

wjuroesafgizkmvplxc
nxvgfoeucjmprikwlsy
glwrfcumhejisvkpox
kjmuvrwclifaxepogs

kjxygrvsodimabecwuzh
ihmxrogvjsaudbkwcezy
ykwovzjimbgashrudxec
jsewruvymigxdzkohabc
icmleouvhyzxdabjsgkwr

wm
wm
miwkg

slcdjyhbnpi
bxqvejnzwtfag

a
a
a
a
a

jxwonlf
jox

wxhytgp
tpgywx
pgwyuxt
pgxyfwta
dpyzwxrg

uygwfkxvane
hsdicwzkrenfgalqx
jaknwfxge

mxs
msk
smx
msoqz
ms

bxfwc
wcbxf
cxfbw
wkcbxf
cfwxbl

om
om
a
n

ikp
wi
iu
biej
i

aokgqimrnlfbdxzjs
tmqxrbislngykoajd
lsgvboxnkjdqrima
orijsalbqnmdkgx

zfpd
biok
bxcqnk
yikx

sulwztbnqkvioafxchemr
qhmcvlexkazuinfrwbsto
mnfutiblxecdqwvkoszhar
bmxzpqrlkafheouvcsntwi
zmchneutwlqoakxbsifrv

equacsmywgpkixfdv
anxwlcuypskjgfivbqr
iyaksgxmvoufehwpqc

twfzqxsnr
aqdsp

yj
dyoje
yj
yhqj
jy

bhwsczaxunjr
xgcswaflpznruqb
ubzwjxkcsyrnmahd

xyvgkjo
kgjvoyx
vxyokjg
govjkxy
jgyxokv

hneolzrtywq
atbs
mtipgd
vts

flu
lfu
ufcl

ad
da
da
ad
tgqad

xfw
bufw

simdubgowtap
wgbsdomaiutp
bsmodwauigtp

qi
o

n
dthokli
nveq

nvdt
cdglvi

sdxfaveuwlqzgircnhyobtkjp
uiqjrdenyagwkzhvpxscobtfl
otiungzxecdbyksrfwqjhvlap
dtbcyojehfwvsnrpkliuzgaxq

zgovqbryphxmjcasuktfl
bwcpqoutvfjykmashrzlgx
iecjhpmxsqytzorflungvabk
vacztqslhmkgoxyrpufdbj
amgfxjqsptdrhovclkzuby

s
s
fs
s

b
b
bm
b

uzc
zbc
cz

clmeswqpuahi
spahequmlwci

hrclmu
nhlcdurfxzi
orshcluv
uheclr

r
k
r

da
ad
ad

hlfmnga
amhgn
hgamn
geahwnm
malgnuh

bhpgndjwecroq
iwxklmtyzsopf

uqosfv
lmbihrxtedk

gyrtdosqjnkalicmwvh
rmlygiwktqcoshdavnj
wimjknqvaglhdsocrty
wkosyadhtqgnmcrvijl

tcpsjerzykogwmqbxinlad
iyzxcgrtqosnbkjaleudmw

dyzsmnkvuegbwlt
iwrcgs
fpswgqc

nawtjrgmpyxlibcshfudvkzq
jiblrgwyhtmxvucqz

vcruatwneqxi
myhgldscjnfb

qakljbptgwryiuxnmoh
mtquwoagbjknxyrlpih
aktuhpgmqxbwjyonril
wqriogjmubxpknvyltha
mpkuhlxgjtanwoiyrqb

ptjrmc
ucfjpm
pjcsm

jofrkgdves
gsdkrivfe
dgsevqkxrf

fvs
vclsd
s
rehtuyi
vfk

eafupodb
eyrdpgbujfoa
ukadqobefiwxp
pfgudeobay

i
p
i
i
i

udlqpozmayew
dsyalxmupwe
lipydaumwe
delyumpwsa

zthwnjiuyosakfpxbclm
lcpyfmisbztuknaxo
xsnlcyfzbpokmatui
zblkpsficaxmyoutn

n
nt
r

meskuctxrwaflovdpzhj
koxzmhvrcydlbwpsjuae

l
pqd
lh
f
h

otzqusgdeyvc
aetugnycqrxz
fhzwtgjqmkyulip

o
o
o
o
o

frejdvqokushlxtcgwani
dmfesnxhwgluorjiayt
zdtispaxeobgjhunwfrl

retxhvnmzsjqpbodif
txnphyomczqfisedrjb

nmwra
rmvawn
arnwm
mnrwa
wnmra

yatwdlfbuc
flydbtwacuo
upaybixdlfvtcrw
tncdbalyswfu

lenfcbjv
jfecnvlb
jlvefcnbr
lvcgqfnaetbji

bkuycegrnpimoxvfzl
xpmklivbuyrcega
lbktipremguxcyv
pbmirkvglyceux
ykibxregmlupcv

rwjeyxizhgt
ywxetqgajzriol
wrigyezjtx
werygtjizx

az
ja

sumxgwqydjhe
whtdyeqguxsp

pncwmq
ansluwbycmxoe
cwdnmj
wqhkmcn

flucitkw
ilucwjktf
lwukticf
lcwifktu
flcubitwk

qlhnvft
hnvtfl

bwsegp
bmpgsew
gpsewb
bgwqesp
pgwbse

ib
mgjsb
bzk
bpmsgj
xb

bjerpdfuo
wjrbf

dkmxn
nkmdx
ndxmk
ndmxk

rnouk
j

m
m
t

mrzq
zpqj
qiz
vzq
zhqi

tkn
nrkt

yj
jsyqp
awyj
yj
jy

xl
dlx
olx

kqdluj
jlkvdqu
qdukjl
qdklpju
ukqldj

wzy
zy
zy
zy
yz

vea
aewk
berga
waek
lyaen

miaoyj
yjvmoiza
yimajo
jymaoi
yjoiam

azcrbo
bzero

fig
jbsagif
gfi

ed
de
ed
de

naezbtk
ubat
batku
baqvtglmo
treba

gozafjiuerdpyscb
lxuhqvmktwn

ysqdjnvwhagmzitpf
vszdapqhijcyln
pcazqibvhsdjyn
rjuaqhspydikznv

qhgazndtiwb
wtbqdahviou
zbqjhwgtai
qiabwntckh
siawbyfpmxtrheql

ej
je

vtqgbimyo
dpwihgamzovlbtf
tmgvboi
emntbvgcio
btgimvoy

kisynhablrjwempcouxqgdftz
xnzrkiufjpmghsqcadotwebyl
rdqezcmgajpwoylnbusxfhkti
shejiunzgdpwxylqmatrbfkoc

lcqenihaktdgurpvmysjfxzow
sldaifypqrkhcumextwvjongz
piyslcrajgfenhmdwuovtkzqx
esxwujvtarpzkoldmyfncihgq
xqfvgosbyruephwjdaztklincm

gyofwzc
jyluxh
dpimyv
rtsbyqa
myluej

iwcedhs
etdich
cwihaeb
ecihby
licvfeh

nwlfiydcst
wnldtsfy
lsnwdtyf
fdwytsln
tfsywdnl

diphkr
ihf

osagihfnuptwcyxqevz
ynzjfrpqoicklmuvbw
nckfilwvquozjpmy
qzwoivcfpyjun

tbjh
noyjvqt
ejzuligpxm
hnj
jysha

fycqbrk
kcndb
czblkq
blfcky

wpgk
putk
lqamxiv
rd
dcw

wmludkgnozfe
aoezwflgkmdun
uzekmdgflnwo
kgflnuzowedm
goezkndmwufl

zmbsxkiwqrh
grjcyiwptzseqxofd
iubrwqlkxzsvn
zaqswibmxr

soealywdutckrfnbigpxz
dntrgfaopclubeykzwsix
gnswdzialxrckybtfuope
fwodizuepagysxnktlbcr
sbdiefgptynkuwzocxalr

tkbry
fyxr
ory

sio
oi
oi
oi

fipqkjmorwcluenvzx
mvowrcjxsuqpfgikenl
cxrmwvkeulpfqnoij

jhqfvlnetysazio
osjqykft
otqyjfuxs

opfrbnw
dbqnhoftpruz
bgeyxrpnfol
royngfpbw

rvmxfogsbcjkpqltyhandi
xafqvhinsodltpkcwy

pqrkslwfijgdo
pymoxvzsg
vpszyog
sehgop

qnmxjglwavzirceu
advwiomgshuqnclzjer
zjqumrcgaiwlevn
kvugjnewclizasqmr
jilzanmtbgfrequvcw

szkftpg
stkpgzf
ptzkgsf

mefgopznxlusjych
iocqfmjlzgyxkpunhe

zemltjrnbso
stbxojkrzgmneqv
zpsjernombt
stjobrmzne

yblxcozivekmqa
epsrvyigmc

xs
s
s

bnmz
fhqg
dxsin
zsm

brpge
tpm
ptm
p
sp

utpsykci
ituxbykcwr

zvxityphsdnuagbmcjwoqrfl
qiuwkplbtczmnvyrsoxhafjgd
qynizxpujhfmbvdcgwleasrot

y
tx
y
h
y

fsqatkdlmpux
yxinuvasdpfml
tsdpxafblzomu
axtfpulsmdg
lxwfqpsajudmk

joanzmhfldbsgwcipkxuqrey
oclnpxwkyqszdjreagiumfhb
ufvoiyxzbrpwdgkcmjsqlanhe

ciqxenmrlh
ctrfhxiqpnl

uoxfbrglpwjkhzstidmcyq
hzslfyjxbdgecairompwkqut
ygolzckfbxdmqtpivjshrwu
kpoxzdlhurgimcsftbqywj

xhgfmplcvwrqneztju
vqbwxtiprfhl
wxaoytfrvhbsldpq

iop
qvxihy
io

ftlgqpix
bpigfl
ufbgiplr

ek
ekb
kclifmge
key
ersk

ytqglpedjriu
jidgqytrulpe
tgjyudrqplei
jeiytpurgdlq

anuqyij
auczniy
yanklsxivou
byacinu
inazyue

tmryzavsnfdoq
ngmjxatvosz

yurzemitvfqnksxpbdlhgjwa
juwvmadgcpytrkfqnbezishlx
wzudtksjrpevnxagfihlbyqm

a
xa
ya
a

gnukiszbvj
nyzbvpksjgmu
sgujnbzeivk
uzgbnksjv

xoq
qox

xlberjukaztfnoycqsgwip
tscfaexqwgkrpouzbylijn
owaruypenjqtbslizcgxkf
jzcuswiqgkrbnpyoxletfa

pxzmb
rzmxpb
bpxznmrc
umbgpzkyox
tzbmxp

grspmxbyhojdizwlevqfkctu
fixgbtedqomczrjhswupkyl
tlsgyobzpmcqkhrdxewfuij

xrhelazymwntkv
lhvamxeyrznw
vrmwalyhnezx

mdbjcizwxh
jpcbzdmxiwh

xakdgjtimlvzcshw
iskdxfjnlhwrmvtcz

yrli
lriz
rli
irol
ilr

lfamcngvojeqyzrkdspwi
slqcznmwdivkrfjopebayg
cwgodjahrqzyfivepknmsl
kyeuvpjmlnftsizrdogqcaw
kjvrcgypqzinwadmoefls

hfmuzdvyibakcwnsor
cezounhbsydwakmif

vsajkf
jvf
vofj

q
v

h
h
i

frgo
ogfr
grfo
rofg
gfro

saomp
mpdosa
posam
mspao
ompsa

qbxi
xbiq
bixoq
axbqi

oix
xoi
ixo

zqkved
ztvydqe
vdekzq

ubg
vgcub

wdotn
akp
hv
ipj

xfcprnv
dfpcrvu
jpvfirc
frphqvwboc
pcvfxr

prx
rlxp
rxp
ixpr
pxr

uaozgkwlbfxyvmi
zbkgoxalfwuvymi
yimugvxbkfolwaz

hwvlbzcauseydgpxrkq
lamwdrzhsbqefkvytxoicu
zktxwlnhysevqbraducj

vjuprhcl
vjfrlpm
jbplvr
pnvljrbfe

cwygrulzbe
elyzgwbu
uygezlwb
yebuglzw
ygzwuble

wbzecjtdrgiyqls
wlsycjgeizdr
zslyderwigcj
dwzcyjgerisl

qkvegzmnobpuc
pogzvqmkbn

cnepixbwhklmqzodvyautjgrfs
oepszjgmlcxkdiartnybwhfvqu
gvcxaintlrweupbzfsmykdjqho
xalwvjenruhpkmgfqtoysbdizc

omrc
tcer
bguaf
vhtk

rtednpkyxgcs
kxlsredupynt
bartsexnypfmqk

lqrabmznuhgjfi
chligsdnbjrazum
jpghuibarnzlkom
bamhwrngzjfliuye
yuignrzsamhlxbjt

aoslfqnjcghb
cnbflgjqaohs
oqshfljbncaxg

av
av
va
dvqa

lghqidarnczwfxyu
ylfnhqwczaiuxdg
wfdglzixauchnqy

pfgknomr
nfegr
vihnlcfdqgy
fgrns
awogfn

kfdntirsmqap
thjdknbugmr
nxfetkrmd

ydabnisroqeghk
gryenohbaiqsk
sqbingvyoakehr
ngbkyfaicsrhqeo

ghxtulq
ntgqxlsz
lgzqtxsi

qvr
mqljsu
wlqu
gpoqytkbzf
evqw

khybpm
rhboj

rmcqdblnto
qlcnmor
rhmzalcsoq
lgxcrmnqovd

wdfkpmalijbncuvr
qhnmikpzaygxwsovej
`

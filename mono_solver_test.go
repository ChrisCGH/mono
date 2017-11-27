package mono

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var english_tri string = `THE : 0 0.99
ING : 0 0.93
AND : 0 0.88
YOU : 0 0.83
ENT : 0 0.82
THA : 0 0.82
FOR : 0 0.81
HER : 0 0.81
HAT : 0 0.80
ION : 0 0.80
VER : 0 0.79
ALL : 0 0.78
ERE : 0 0.78
ERS : 0 0.78
EST : 0 0.78
ETH : 0 0.78
NTH : 0 0.78
TER : 0 0.78
INT : 0 0.77
IST : 0 0.77
STH : 0 0.77
ARE : 0 0.76
ATI : 0 0.76
EAR : 0 0.76
ITH : 0 0.76
ONE : 0 0.76
OUR : 0 0.76
REA : 0 0.76
THI : 0 0.76
TIO : 0 0.76
GHT : 0 0.75
MAN : 0 0.75
STO : 0 0.75
WIT : 0 0.75
ATE : 0 0.74
EVE : 0 0.74
OME : 0 0.74
OUT : 0 0.74
TIN : 0 0.74
TTH : 0 0.74
ANT : 0 0.73
AST : 0 0.73
ESS : 0 0.73
HEN : 0 0.73
HIN : 0 0.73
IGH : 0 0.73
IVE : 0 0.73
NOT : 0 0.73
OTH : 0 0.73
SAN : 0 0.73
SIN : 0 0.73
AVE : 0 0.72
CAN : 0 0.72
COM : 0 0.72
DTH : 0 0.72
HES : 0 0.72
HIS : 0 0.72
ILL : 0 0.72
INE : 0 0.72
ONT : 0 0.72
OUS : 0 0.72
PER : 0 0.72
RES : 0 0.72
STA : 0 0.72
ERI : 0 0.71
ERT : 0 0.71
NCE : 0 0.71
NIN : 0 0.71
ONS : 0 0.71
ORE : 0 0.71
ORT : 0 0.71
OUN : 0 0.71
STE : 0 0.71
UND : 0 0.71
WHE : 0 0.71
CON : 0 0.70
DIN : 0 0.70
EAN : 0 0.70
ECO : 0 0.70
ESA : 0 0.70
HAN : 0 0.70
HEM : 0 0.70
ITI : 0 0.70
LES : 0 0.70
NGT : 0 0.70
OFT : 0 0.70
OVE : 0 0.70
PRO : 0 0.70
RIN : 0 0.70
RTH : 0 0.70
THO : 0 0.70
WHO : 0 0.70
WOR : 0 0.70`

var english_tri_with_spaces string = `{TH : 53839 11.5909
THE : 44018 11.3895
HE{ : 43556 11.379
{AN : 25982 10.8623
AND : 24896 10.8196
ED{ : 24785 10.8152
ND{ : 24499 10.8036
NG{ : 19492 10.5749
ING : 19229 10.5613
{TO : 17114 10.4448
TO{ : 16467 10.4063
{HE : 16183 10.3889
ER{ : 15868 10.3692
AT{ : 14960 10.3103
{OF : 14366 10.2698
{HI : 13792 10.229
OF{ : 13733 10.2247
D{T : 12925 10.1641
IS{ : 12727 10.1487
AS{ : 12124 10.1001
HER : 11463 10.0441
{IN : 11434 10.0415
RE{ : 11078 10.0099
E{A : 10928 9.99625
{HA : 10744 9.97927
IN{ : 10652 9.97067
{WH : 10136 9.92102
HAT : 10132 9.92063
E{T : 10116 9.91904
HIS : 9734 9.88055
ON{ : 9645 9.87137
{WA : 9296 9.83451
{A{ : 9110 9.8143
T{T : 9107 9.81397
THA : 9066 9.80946
{BE : 9025 9.80493
{CO : 8936 9.79501
E{S : 8598 9.75646
LY{ : 8476 9.74217
S{A : 8272 9.7178
{WI : 8262 9.71659
D{A : 8261 9.71647
E{W : 8147 9.70258
{NO : 8126 9.7
ES{ : 8013 9.68599
N{T : 7987 9.68274
EN{ : 7915 9.67369
ERE : 7909 9.67293
E{H : 7512 9.62143
LL{ : 7310 9.59417
ENT : 7273 9.5891
NT{ : 7059 9.55923
{ON : 7013 9.55269
UT{ : 6964 9.54568
D{H : 6850 9.52918
WAS : 6804 9.52244
{RE : 6759 9.5158
TH{ : 6695 9.50629
WIT : 6601 9.49215
AD{ : 6577 9.48851
OR{ : 6576 9.48835
ITH : 6475 9.47288
E{O : 6437 9.46699
FOR : 6392 9.45997
S{T : 6353 9.45385
CE{ : 6214 9.43173
NOT : 6056 9.40598
{SO : 6000 9.39669
{SH : 5999 9.39652
NCE : 5936 9.38596
{SA : 5870 9.37478
AN{ : 5849 9.3712
LE{ : 5844 9.37034
TER : 5810 9.36451
{FO : 5803 9.3633
T{A : 5729 9.35047
F{T : 5718 9.34855
ION : 5705 9.34627
E{C : 5636 9.3341
ME{ : 5629 9.33286
ALL : 5611 9.32966
THI : 5607 9.32894
IT{ : 5449 9.30036
N{A : 5337 9.27959
{IT : 5322 9.27678
{MA : 5315 9.27546
T{H : 5209 9.25531
ESS : 5168 9.24741
{AS : 5150 9.24392
YOU : 5112 9.23652
{WE : 5098 9.23377
{PR : 5056 9.2255
VER : 5047 9.22372
OT{ : 5035 9.22134
HIM : 5023 9.21895
HAD : 4969 9.20815
RIN : 4926 9.19945
LD{ : 4922 9.19864
{FR : 4883 9.19069
CH{ : 4876 9.18925
E{F : 4853 9.18452
{YO : 4849 9.1837
ERS : 4838 9.18143
{AL : 4827 9.17915
TIO : 4762 9.16559
SE{ : 4740 9.16096
ID{ : 4717 9.1561
OW{ : 4676 9.14737
SHE : 4663 9.14459
VE{ : 4648 9.14136
S{O : 4625 9.1364
ST{ : 4608 9.13272
{AT : 4602 9.13142
D{S : 4556 9.12137
{ST : 4531 9.11587
HIN : 4531 9.11587
E{I : 4508 9.11078
{SE : 4448 9.09738
{MO : 4429 9.0931
S{S : 4416 9.09016
O{T : 4380 9.08198
G{T : 4372 9.08015
S{W : 4358 9.07694
ONE : 4341 9.07303
NE{ : 4296 9.06261
IM{ : 4272 9.05701
R{T : 4259 9.05396
D{B : 4254 9.05279
Y{A : 4253 9.05255
T{W : 4232 9.0476
T{I : 4219 9.04452
E{B : 4172 9.03332
{CA : 4171 9.03308
HOU : 4165 9.03164
OUN : 4138 9.02514
{BU : 4135 9.02441
T{O : 4118 9.02029
RES : 4091 9.01372
RS{ : 4072 9.00906`

var english_tet string = `THAT : 5307 8.92778
THER : 5075 8.88308
WITH : 4325 8.72317
NTHE : 3957 8.63424
DTHE : 3369 8.47337
HERE : 3356 8.46951
TION : 2872 8.31377
OTHE : 2668 8.24009
SAID : 2579 8.20616
OULD : 2565 8.20071
TTHE : 2427 8.14541
THEM : 2329 8.10419
THES : 2319 8.0999
EAND : 2318 8.09946
THEC : 2287 8.086
ETHE : 2239 8.06478
INTH : 2215 8.05401
OFTH : 2192 8.04357
ANDT : 2169 8.03302
SAND : 2153 8.02562
IGHT : 2056 7.97952
HAVE : 2054 7.97855
ENCE : 2039 7.97121
FTHE : 1972 7.9378
ANDS : 1961 7.93221
THIS : 1934 7.91835
EVER : 1920 7.91108
INGT : 1906 7.90376
THIN : 1863 7.88094
OMBE : 1825 7.86033
NDTH : 1819 7.85704
DOMB : 1798 7.84543
MBEY : 1792 7.84209
TAND : 1711 7.79584
EDTH : 1694 7.78585
ONTH : 1679 7.77696
OUGH : 1667 7.76978
TAIN : 1665 7.76858
DAND : 1658 7.76437
INGA : 1621 7.7418
STHE : 1591 7.72312
VERY : 1498 7.66289
TTLE : 1484 7.6535
TOTH : 1470 7.64402
EDTO : 1435 7.61992
ATTH : 1417 7.6073
HING : 1409 7.60164
TING : 1398 7.5938
EWAS : 1382 7.58229
INGH : 1378 7.57939
RENC : 1372 7.57503
ATIO : 1332 7.54544
RTHE : 1332 7.54544
CAPT : 1322 7.5379
APTA : 1311 7.52955
NGTH : 1309 7.52802
PTAI : 1308 7.52726
WHEN : 1303 7.52343
ANDW : 1299 7.52035
HICH : 1293 7.51572
WHIC : 1291 7.51417
UGHT : 1290 7.5134
HEHA : 1285 7.50952
HAND : 1283 7.50796
THOU : 1276 7.50249
KING : 1273 7.50013
ATHE : 1257 7.48748
OREN : 1253 7.4843
SELF : 1250 7.4819
MENT : 1242 7.47548
THEW : 1230 7.46577
YAND : 1228 7.46414
THED : 1212 7.45103
LORE : 1208 7.44772
WERE : 1208 7.44772
RAND : 1207 7.44689
FLOR : 1204 7.44441
FROM : 1203 7.44358
ANDA : 1201 7.44191
EFOR : 1198 7.43941
NING : 1191 7.43355
WHAT : 1189 7.43187
HECA : 1178 7.42257
ERTH : 1176 7.42087
ANDH : 1172 7.41747
LOOK : 1170 7.41576
SOME : 1170 7.41576
WOUL : 1154 7.40199
THAN : 1150 7.39852
RDOM : 1135 7.38539
ITHA : 1127 7.37832
UPON : 1125 7.37654
EHAD : 1109 7.36221
YOUR : 1105 7.3586
NAND : 1086 7.34126
THEY : 1085 7.34034
MRDO : 1084 7.33941
HATI : 1082 7.33757
HTHE : 1077 7.33294
INGI : 1069 7.32548
HEWA : 1068 7.32454
HERS : 1066 7.32267
MISS : 1063 7.31985
RING : 1054 7.31135
NHIS : 1048 7.30564
INGO : 1036 7.29412
HEMA : 1011 7.2697
BEEN : 1011 7.2697
TIME : 1010 7.26871
ERED : 999 7.25776
KNOW : 995 7.25374
OUND : 992 7.25072
TURN : 987 7.24567
THEP : 983 7.24161
THEN : 971 7.22933
ETHA : 969 7.22727
LITT : 968 7.22623
ANDI : 966 7.22416
INGS : 963 7.22106
REAT : 963 7.22106
NDER : 962 7.22002
FORT : 961 7.21898
EDAN : 961 7.21898
ITTL : 950 7.20746
THEB : 948 7.20536
ECAP : 941 7.19794
THEL : 939 7.19582
LING : 937 7.19368
DHER : 935 7.19155
THEH : 928 7.18403
HATH : 924 7.17971
EYOU : 913 7.16774
YTHE : 909 7.16335
TWAS : 909 7.16335
NHER : 906 7.16004
INTO : 900 7.1534
ERAN : 889 7.1411
STAN : 884 7.13546
ERIN : 882 7.13319
ABLE : 871 7.12064
ORTH : 870 7.11949
TOBE : 863 7.11141
THEI : 861 7.1091
FORE : 859 7.10677
HOUG : 855 7.1021
THEF : 850 7.09624
EDIN : 847 7.0927
ANCE : 842 7.08678
DING : 839 7.08321
GTHE : 837 7.08082
COUL : 837 7.08082
DTHA : 836 7.07963
MORE : 831 7.07363
EDHI : 824 7.06517
OFHI : 823 7.06396
PRES : 818 7.05786
NGTO : 808 7.04556
HESA : 806 7.04308
GAIN : 806 7.04308
TTER : 799 7.03436
NESS : 799 7.03436
HECO : 798 7.03311
EDIT : 796 7.0306
LIKE : 795 7.02934
HISH : 792 7.02556
THET : 789 7.02177
HERA : 781 7.01158
RESS : 779 7.00901
ERES : 776 7.00515`

var english_tet_with_spaces string = `{THE : 36126 11.1919
THE{ : 31344 11.0499
AND{ : 21758 10.6849
{AND : 19717 10.5864
ING{ : 17874 10.4883
{TO{ : 14673 10.2909
{OF{ : 13033 10.1724
HAT{ : 9581 9.86471
HIS{ : 9102 9.81342
{HE{ : 8361 9.72851
{THA : 8122 9.6995
D{TH : 7967 9.68024
{IN{ : 7713 9.64783
THAT : 7617 9.63531
{HIS : 7032 9.5554
WAS{ : 6746 9.51388
HER{ : 6746 9.51388
{WAS : 6377 9.45762
T{TH : 6355 9.45417
WITH : 6073 9.40878
N{TH : 6067 9.40779
{WIT : 6003 9.39719
ED{T : 5976 9.39268
ITH{ : 5300 9.27263
F{TH : 5256 9.2643
THER : 5222 9.25781
E{TH : 5168 9.24741
OF{T : 5031 9.22055
HE{S : 4917 9.19763
E{AN : 4914 9.19702
HAD{ : 4895 9.19314
{NOT : 4841 9.18205
ERE{ : 4837 9.18122
{HER : 4761 9.16538
{HIM : 4601 9.1312
{HAD : 4587 9.12815
{YOU : 4547 9.11939
NOT{ : 4472 9.10276
{FOR : 4410 9.0888
ND{T : 4407 9.08812
TION : 4324 9.06911
NG{T : 4291 9.06145
{IT{ : 4275 9.05771`

var english_ngraph string = `PRINCEA : 1133 7.91851
RINCEAN : 1104 7.89258
INCEAND : 1100 7.88895
NATASHA : 1092 7.88165
CEANDRE : 1082 7.87245
NCEANDR : 1077 7.86782
PRINCES : 1026 7.81931
EANDREW : 965 7.75801
THATTHE : 964 7.75697
THOUGHT : 942 7.73389
HIMSELF : 903 7.69161
RINCESS : 869 7.65323
FROMTHE : 858 7.64049
WITHTHE : 840 7.61929
THECOUN : 764 7.52445
HECOUNT : 700 7.43696
EPRINCE : 692 7.42547
HADBEEN : 691 7.42402
OULDNOT : 663 7.38266
THEFREN : 657 7.37357
HEFRENC : 657 7.37357
UNDERST : 642 7.35047
EFRENCH : 634 7.33793
NICHOLA : 633 7.33635
OFFICER : 610 7.29934
SOMETHI : 608 7.29606
EMPEROR : 606 7.29276
THEEMPE : 605 7.29111
OMETHIN : 605 7.29111
EEMPERO : 605 7.29111
HEEMPER : 601 7.28448
GENERAL : 601 7.28448
DPRINCE : 600 7.28281
NAPOLEO : 593 7.27108
WITHOUT : 591 7.2677
RUSSIAN : 574 7.23851
METHING : 566 7.22448
ICHOLAS : 566 7.22448
THESAME : 553 7.20124
COMMAND : 552 7.19943
ANOTHER : 550 7.1958
APOLEON : 544 7.18483
EXPRESS : 533 7.1644
SOLDIER : 532 7.16253
COULDNO : 520 7.13971
INCESSM : 518 7.13586
NCESSMA : 514 7.12811
CESSMAR : 512 7.12421
ESSMARY : 504 7.10846
ANDTHAT : 503 7.10647
COUNTES : 489 7.07825
KUTUZOV : 487 7.07415
BECAUSE : 456 7.00838
OUNTESS : 440 6.97266
NDERSTA : 440 6.97266
DERSTAN : 439 6.97038
ETHOUGH : 435 6.96123
NOTHING : 434 6.95893
HEPRINC : 433 6.95662
THEPRIN : 432 6.95431
SUDDENL : 432 6.95431
WITHHIS : 431 6.95199
EVERYTH : 431 6.95199
VERYTHI : 425 6.93797
ERYTHIN : 423 6.93326
EVIDENT : 419 6.92375
INTOTHE : 418 6.92136
ERSTAND : 405 6.88977
SAIDTHE : 396 6.8673
UDDENLY : 390 6.85203
LOOKING : 389 6.84946
PRESSIO : 386 6.84172
RYTHING : 385 6.83913
INGTHAT : 385 6.83913
POSITIO : 383 6.83392
THEOTHE : 381 6.82868
OMMANDE : 381 6.82868
EDIDNOT : 376 6.81547
DENISOV : 375 6.81281
SANDTHE : 373 6.80746
THEFIRS : 371 6.80209
CONTINU : 371 6.80209
SEEMEDT : 370 6.79939
THROUGH : 366 6.78852
CHAPTER : 366 6.78852
RESSION : 365 6.78578
POSSIBL : 365 6.78578
QUESTIO : 359 6.76921
ERUSSIA : 358 6.76642
FEELING : 357 6.76362
XPRESSI : 352 6.74951
INGWITH : 351 6.74667
VIDENTL : 350 6.74382
ECOUNTE : 349 6.74096
HEFIRST : 348 6.73809
OLDIERS : 347 6.73521
OUGHTHE : 345 6.72943
HEOTHER : 343 6.72361
DPIERRE : 342 6.72069
OSITION : 338 6.70893`

var english_ngraph_with_spaces string = `{THAT{ : 6710 9.50853
{WITH{ : 4843 9.18246
{OF{TH : 4605 9.13207
N{THE{ : 3995 8.98997
D{THE{ : 3962 8.98168
OF{THE : 3841 8.95066
F{THE{ : 3578 8.87973
{AND{T : 3241 8.78081
T{THE{ : 3233 8.77834
S{AND{ : 3191 8.76526
{THE{S : 3066 8.7253
E{AND{ : 2892 8.66688
AND{TH : 2858 8.65505
{THE{C : 2842 8.64944
PRINCE : 2712 8.60261
TO{THE : 2554 8.54259
{TO{TH : 2489 8.51681
{IN{TH : 2477 8.51198
{SAID{ : 2458 8.50428
ED{TO{ : 2420 8.4887
O{THE{ : 2417 8.48745
{PRINC : 2391 8.47664
AT{THE : 2321 8.44693
IN{THE : 2314 8.44391
{FROM{ : 2292 8.43435
ND{THE : 2278 8.42823
D{AND{ : 2229 8.40648
ING{TH : 2198 8.39248
{THE{F : 2179 8.38379
{WERE{ : 2044 8.31984
D{NOT{ : 2024 8.31
ED{THE : 2014 8.30505
E{WAS{ : 1935 8.26503
WHICH{ : 1900 8.24678
{THE{P : 1881 8.23673
THING{ : 1876 8.23407
{AND{S : 1855 8.22281
{WHAT{ : 1841 8.21524
{THE{R : 1831 8.20979
E{THE{ : 1823 8.20541
THE{CO : 1820 8.20376
E{HAD{ : 1772 8.17704
{THEY{ : 1768 8.17478
NG{THE : 1752 8.16569
PIERRE : 1734 8.15536
RINCE{ : 1731 8.15363
{THIS{ : 1709 8.14084
{HAVE{ : 1699 8.13497
{WHICH : 1697 8.13379
S{THE{ : 1677 8.12193
{AND{H : 1622 8.08859
{AND{A : 1612 8.0824
Y{AND{ : 1610 8.08116
{THE{M : 1609 8.08054
IERRE{ : 1567 8.05409
{PIERR : 1565 8.05281
R{THE{ : 1494 8.00638
{AT{TH : 1458 7.98199
{THE{D : 1449 7.9758
{AND{W : 1448 7.97511
{THE{E : 1446 7.97373
T{AND{ : 1443 7.97165
CE{AND : 1417 7.95347
ON{THE : 1409 7.94781
R{AND{ : 1391 7.93495
{THE{W : 1385 7.93063
{THERE : 1375 7.92338
OTHER{ : 1371 7.92047
{THE{B : 1355 7.90873
THERE{ : 1352 7.90651
G{THE{ : 1346 7.90206
HE{HAD : 1338 7.8961
THOUGH : 1336 7.89461
{THE{O : 1336 7.89461
HAT{TH : 1329 7.88935
ING{TO : 1321 7.88332
T{WAS{ : 1319 7.8818
N{AND{ : 1318 7.88104
THAT{T : 1315 7.87876
{ON{TH : 1296 7.86421
D{THAT : 1295 7.86344
THEIR{ : 1283 7.85413
ED{AND : 1280 7.85179
WOULD{ : 1264 7.83921
NG{TO{ : 1252 7.82967
{TO{HI : 1251 7.82887
{THE{T : 1240 7.82004
ATION{ : 1230 7.81194
Y{THE{ : 1225 7.80787
{BEEN{ : 1220 7.80378
NATASH : 1213 7.79802
D{HIS{ : 1202 7.78891
{THEM{ : 1196 7.78391
NCE{AN : 1193 7.7814
{THOUG : 1192 7.78056
{HE{HA : 1188 7.7772
{THE{A : 1187 7.77636
{THEIR : 1172 7.76364
{THE{H : 1163 7.75593
{WOULD : 1152 7.74643`

func writeTheTrigramFile(t *testing.T, filename string) {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	defer f.Close()
	if err != nil {
		t.Fatal("Failed to create trigram file")
	}
	fmt.Fprint(f, english_tri)
}

func writeTheTrigramWithSpacesFile(t *testing.T, filename string) {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	defer f.Close()
	if err != nil {
		t.Fatal("Failed to create trigram file")
	}
	fmt.Fprint(f, english_tri_with_spaces)
}

func writeTheTetragramFile(t *testing.T, filename string) {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	defer f.Close()
	if err != nil {
		t.Fatal("Failed to create tetragram file")
	}
	fmt.Fprint(f, english_tet)
}

func writeTheTetragramWithSpacesFile(t *testing.T, filename string) {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	defer f.Close()
	if err != nil {
		t.Fatal("Failed to create tetragram file")
	}
	fmt.Fprint(f, english_tet_with_spaces)
}

func writeTheNgraphFile(t *testing.T, filename string) {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	defer f.Close()
	if err != nil {
		t.Fatal("Failed to create ngraph file")
	}
	fmt.Fprint(f, english_ngraph)
}

func writeTheNgraphWithSpacesFile(t *testing.T, filename string) {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	defer f.Close()
	if err != nil {
		t.Fatal("Failed to create ngraph file")
	}
	fmt.Fprint(f, english_ngraph_with_spaces)
}

func TestNewMono_solver(t *testing.T) {
	ms := NewMono_Solver()
	if ms.max_top_ != 1 {
		t.Error("ms.max_top_ should be 1")
	}
}

func TestSet_cipher_text(t *testing.T) {
	ms := NewMono_Solver()

	test_tri_file := "test.tri"
	writeTheTrigramFile(t, test_tri_file)
	defer os.Remove(test_tri_file)
	(&ms).Set_trigraph_scoring(test_tri_file)
	(&ms).Set_cipher_text("ABC DEF")
	if ms.ciphertext_ != "ABCDEF" {
		t.Errorf("ms.ciphertext_ should be %s, but actually is %s\n", "ABCDEF", ms.ciphertext_)
	}

	test_tri_spaces_file := "test_spaces.tri"
	writeTheTrigramWithSpacesFile(t, test_tri_spaces_file)
	defer os.Remove(test_tri_spaces_file)
	fmt.Println("Trigraph")
	(&ms).Set_trigraph_scoring(test_tri_spaces_file)
	(&ms).Set_cipher_text("ABC DEF")
	if ms.ciphertext_ != "ABC{DEF" {
		t.Errorf("ms.ciphertext_ should be %s, but actually is %s\n", "ABC{DEF", ms.ciphertext_)
	}

	test_tet_file := "test.tet"
	writeTheTetragramFile(t, test_tet_file)
	defer os.Remove(test_tet_file)
	fmt.Println("Tetragraph")
	(&ms).Set_tetragraph_scoring(test_tet_file)
	(&ms).Set_cipher_text("ABC DEF")
	if ms.ciphertext_ != "ABCDEF" {
		t.Errorf("ms.ciphertext_ should be %s, but actually is %s\n", "ABCDEF", ms.ciphertext_)
	}

	test_tet_spaces_file := "test_spaces.tet"
	writeTheTetragramWithSpacesFile(t, test_tet_spaces_file)
	defer os.Remove(test_tet_spaces_file)
	(&ms).Set_tetragraph_scoring(test_tet_spaces_file)
	(&ms).Set_cipher_text("ABC DEF")
	if ms.ciphertext_ != "ABC{DEF" {
		t.Errorf("ms.ciphertext_ should be %s, but actually is %s\n", "ABC{DEF", ms.ciphertext_)
	}

	test_ngraph_file := "test.ngraph"
	writeTheNgraphFile(t, test_ngraph_file)
	defer os.Remove(test_ngraph_file)
	(&ms).Set_ngraph_scoring(test_ngraph_file)
	(&ms).Set_cipher_text("ABC DEF")
	if ms.ciphertext_ != "ABCDEF" {
		t.Errorf("ms.ciphertext_ should be %s, but actually is %s\n", "ABCDEF", ms.ciphertext_)
	}

	test_ngraph_spaces_file := "test_spaces.ngraph"
	writeTheNgraphWithSpacesFile(t, test_ngraph_spaces_file)
	defer os.Remove(test_ngraph_spaces_file)
	(&ms).Set_ngraph_scoring(test_ngraph_spaces_file)
	(&ms).Set_cipher_text("ABC DEF")
	if ms.ciphertext_ != "ABC{DEF" {
		t.Errorf("ms.ciphertext_ should be %s, but actually is %s\n", "ABC{DEF", ms.ciphertext_)
	}
}

func Test_setters(t *testing.T) {
	ms := NewMono_Solver()
	(&ms).Set_verbose()
	if !ms.verbose_ {
		t.Error("ms.verbose_ should be true")
	}

	mi := map[int]int{20000000: 20000000, -1: 0}
	for k, v := range mi {
		(&ms).Set_max_iterations(k)
		if ms.max_iterations_ != v {
			t.Errorf("ms.max_iterations_ should be %d, but actually is %d\n", v, ms.max_iterations_)
		}
	}

	to := map[int]int{120: 120, -1: 0}
	for k, v := range to {
		(&ms).Set_timeout(k)
		if ms.timeout_ != v {
			t.Errorf("ms.timeout_ should be %d, but actually is %d\n", v, ms.timeout_)
		}
	}

	f := NewFixed_Key()
	(&f).Set(byte('e'), byte('J'))
	(&ms).Set_fixed(f)

	if !ms.fixed_.Is_set(byte('e')) {
		t.Errorf("e should be set in ms.fixed_\n")
	}
	if ms.fixed_.Get_pt(byte('J')) != byte('e') {
		t.Errorf("J should be set to e in ms.fixed_\n")
	}
}

func TestIs_time_to_stop(t *testing.T) {
	ms := NewMono_Solver()
	if ms.is_time_to_stop() {
		t.Error("ms.is_time_to_stop() should return false")
	}

	test_tri_file := "test.tri"
	writeTheTrigramFile(t, test_tri_file)
	defer os.Remove(test_tri_file)
	(&ms).Set_trigraph_scoring(test_tri_file)

	ms.max_iterations_ = 10
	if ms.is_time_to_stop() {
		t.Error("ms.is_time_to_stop() should return false")
	}

	for i := 0; i < 10; i++ {
		_ = ms.scorer_.Score("abc", false)
	}
	if ms.is_time_to_stop() {
		t.Error("ms.is_time_to_stop() should return false")
	}

	_ = ms.scorer_.Score("abc", false)
	if !ms.is_time_to_stop() {
		t.Error("ms.is_time_to_stop() should return true")
	}

	ms.max_iterations_ = 100
	ms.timeout_ = 10
	if ms.is_time_to_stop() {
		t.Error("ms.is_time_to_stop() should return false")
	}

	ms.elapsed_ = time.Duration(5 * time.Second)
	if ms.is_time_to_stop() {
		t.Error("ms.is_time_to_stop() should return false")
	}

	ms.elapsed_ = time.Duration(11 * time.Second)
	if !ms.is_time_to_stop() {
		t.Error("ms.is_time_to_stop() should return true")
	}
}

func TestSolve(t *testing.T) {
	ms := NewMono_Solver()

	test_tri_spaces_file := "test_spaces.tri"
	writeTheTrigramWithSpacesFile(t, test_tri_spaces_file)
	defer os.Remove(test_tri_spaces_file)
	(&ms).Set_trigraph_scoring(test_tri_spaces_file)
	(&ms).Set_cipher_text("ABC DEF")
	if ms.ciphertext_ != "ABC{DEF" {
		t.Errorf("ms.ciphertext_ should be %s, but actually is %s\n", "ABC{DEF", ms.ciphertext_)
	}
	(&ms).Set_max_iterations(10)
	(&ms).Set_verbose()

	_ = (&ms).Solve()

	ct := `BT JPX RMLX PCUV AMLX ICVJP IBTWXVR CI M LMT'R PMTN, MTN
YVCJX CDXV MWMBTRJ JPX AMTNGXRJBAH UQCT JPX QGMRJXV CI JPX
YMGG CI JPX HBTW'R QMGMAX; MTN JPX HBTW RMY JPX QMVJ CI JPX
PMTN JPMJ YVCJX. JPXT JPX HBTW'R ACUTJXTMTAX YMR APMTWXN,
MTN PBR JPCUWPJR JVCUFGXN PBL, RC JPMJ JPX SCBTJR CI PBR
GCBTR YXVX GCCRXN, MTN PBR HTXXR RLCJX CTX MWMBTRJ
MTCJPXV. JPX HBTW AVBXN MGCUN JC FVBTW BT JPX MRJVCGCWXVR,
JPX APMGNXMTR, MTN JPX RCCJPRMEXVR. MTN JPX HBTW RQMHX,
MTN RMBN JC JPX YBRX LXT CI FMFEGCT, YPCRCXDXV RPMGG VXMN
JPBR YVBJBTW, MTN RPCY LX JPX BTJXVQVXJMJBCT JPXVXCI,
RPMGG FX AGCJPXN YBJP RAMVGXJ, MTN PMDX M APMBT CI WCGN
MFCUJ PBR TXAH, MTN RPMGG FX JPX JPBVN VUGXV BT JPX
HBTWNCL. JPXT AMLX BT MGG JPX HBTW'R YBRX LXT; FUJ JPXE
ACUGN TCJ VXMN JPX YVBJBTW, TCV LMHX HTCYT JC JPX HBTW JPX
BTJXVQVXJMJBCT JPXVXCI. JPXT YMR HBTW FXGRPMOOMV WVXMJGE
JVCUFGXN, MTN PBR ACUTJXTMTAX YMR APMTWXN BT PBL, MTN PBR
GCVNR YXVX MRJCTBRPXN. TCY JPX KUXXT, FE VXMRCT CI JPX
YCVNR CI JPX HBTW MTN PBR GCVNR, AMLX BTJC JPX FMTKUXJ
PCURX; MTN JPX KUXXT RQMHX MTN RMBN, C HBTW, GBDX ICVXDXV;
GXJ TCJ JPE JPCUWPJR JVCUFGX JPXX, TCV GXJ JPE ACUTJXTMTAX
FX APMTWXN; JPXVX BR M LMT BT JPE HBTWNCL, BT YPCL BR JPX
RQBVBJ CI JPX PCGE WCNR; MTN BT JPX NMER CI JPE IMJPXV
GBWPJ MTN UTNXVRJMTNBTW MTN YBRNCL, GBHX JPX YBRNCL CI JPX
WCNR, YMR ICUTN BT PBL; YPCL JPX HBTW TXFUAPMNTXOOMV JPE
IMJPXV, JPX HBTW, B RME, JPE IMJPXV, LMNX LMRJXV CI JPX
LMWBABMTR, MRJVCGCWXVR, APMGNXMTR, MTN RCCJPRMEXVR;
ICVMRLUAP MR MT XZAXGGXTJ RQBVBJ, MTN HTCYGXNWX, MTN
UTNXVRJMTNBTW, BTJXVQVXJBTW CI NVXMLR, MTN RPCYBTW CI PMVN
RXTJXTAXR, MTN NBRRCGDBTW CI NCUFJR, YXVX ICUTN BT JPX
RMLX NMTBXG, YPCL JPX HBTW TMLXN FXGJXRPMOOMV; TCY GXJ
NMTBXG FX AMGGXN, MTN PX YBGG RPCY JPX BTJXVQVXJMJBCT. JPX
IBVRJ ACNXYCVN BR CJPXGGC.`
	(&ms).Set_cipher_text(ct)
	(&ms).Set_max_iterations(100000)
	(&ms).verbose_ = true
	_ = (&ms).Solve()
	fmt.Printf("Solution = %s\n", ms.Solution())
	fmt.Printf("Key = %s\n", ms.Key())
	fmt.Printf("Score = %d\n", ms.Score())
	fmt.Printf("Iterations = %d\n", ms.Iterations())
	fmt.Printf("Elapsed = %s\n", ms.Elapsed())
}

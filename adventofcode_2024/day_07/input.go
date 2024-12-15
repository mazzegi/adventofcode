package day_07

const input = `
88045337: 4 7 80 453 37
1607806: 4 348 49 228 92 1 14
327264767: 58 8 992 711
11278498: 658 7 169 25 374 95 1
49032: 9 3 5 36 588 9
7218: 6 599 4 2 9 2
515788378563: 534 195 972 49 26 4
95881986: 8 78 8 46 4 986 756 3 6
24033: 77 3 5 3 5 4 33 1
17081838: 1 588 29 827 8
96767: 8 4 91 65 576
277518129: 6 968 74 6 550 347 7
36228: 36 130 81 7 9
698187755: 997 4 1 7 756
34951709: 7 6 7 2 8 765 9 97 1 9 1 7
25662399: 1 60 7 6 4 23 9 6
57179207: 14 1 50 2 1 463 89
320286213: 5 6 99 75 95 5 3 3 1 6 2 7
25533792: 557 53 1 3 40 55 4 4 9 8
582217174: 350 467 90 8 5 166 89
1435524: 31 4 41 52 2
89048732059840: 908 846 69 69 71 2 40
6437556: 34 67 9 6 1 5 1 3 2 9 179
183858813930: 9 702 1 72 3 97 30
374759592: 73 8 5 1 3 571 2 7 1 59 1
383620: 220 2 1 9 192 4 1
14632564: 1 4 110 5 659 4 275 9 4
1206592: 4 60 43 61 9
98225: 38 50 75 7 6 2
21609048657: 81 9 686 139 7 5 2 4
11657: 20 2 286 25 35 2
409975111656: 6 8 949 79 90 869 789
5078158: 5 6 6 92 83 1 4 1 8 4 60 1
628480082978: 3 637 4 491 5 8 2 9 7 4 4
188045: 1 31 112 9 87 6 7 28
5021: 11 95 369 31 72 6 3 3 1
4055703935: 681 215 28 277 672 7
518: 5 37 4 94
168: 7 1 6 2 2
668401968: 267 35 17 500 272 6 8
176765: 8 4 6 982 5
5546996: 562 97 951 6 8 8 351
2536267308: 9 54 7 9 5 2 53 3 1 6 5 5
18542520: 73 847 31 3 4 2 9 5 8 81
207334: 3 16 479 9 407
152492251: 58 53 54 969 481 343
10692549687: 89 4 861 6 7 6 2 2 1 969
10297: 6 5 964 4 6
1560333: 2 8 7 727 26 4 2 8 281 4
72936: 92 9 60 9 3 4 3 64 4 1 4
2597: 470 928 3 834 362
12317609: 79 6 263 551 1
208768594: 9 5 9 68 799 6 233 594
50361046: 281 4 8 8 56 86 50 7
5309045: 1 6 9 9 2 5 203 2 2 1 54 4
7515164: 2 42 93 1 459 4 4 412
313208398725: 6 2 931 16 3 7 711 477
300048: 668 21 63 7 1 57
34246606707: 7 8 7 45 312 967 3
35504479: 70 1 50 13 26 5 79
633051312: 5 3 1 767 95 543 7 185
5880745940: 13 7 98 3 708 37 940
7988276: 5 83 383 3 79 25 403
614299194: 86 5 21 71 96
928219: 88 81 401 1 7 1
322951496024: 53 825 8 9 2 8 947 75
35824: 60 38 91 4 153
1390109: 888 2 11 19 24 71
35266523923: 48 119 103 49 126
2174: 7 19 6 551 5 368 5 2 2
274758: 93 295 40 8
5952821: 744 8 77 8 43
39024913: 3 3 9 1 2 48 916
24224603117: 466 3 79 2 96 685 75
60015: 892 1 7 8 185 8 4 2 993
6157681: 5 23 509 6 11 36 7 2 6 5
20592: 74 67 4 740 20
747889: 988 56 9 184 8 78 25
4061704: 431 94 85 91 829
60765552474: 6 8 82 3 7 8 2 3 648 473
53549: 742 9 8 4 121
223155: 5 551 2 3 6 229 9 8 3 4 3
346872780: 6 6 4 7 976 7 487 4 9 6 3
228105: 3 1 370 609 315 24
27022565: 3 9 6 6 2 1 7 60 22 4 161
5025: 67 3 25
24565051: 7 17 9 2 6 7 3 4 6 525 2
109406696: 5 4 92 3 9 664 77
161291: 9 28 640 9 2
82144310: 865 6 8 766 122
1035450: 74 42 7 3 885
663707: 117 448 4 94 705
730448: 873 3 410 71 8
1370109: 9 703 27 618 3 1
14328144: 4 57 1 446 9 53 517 5 1
84131040: 8 16 4 3 7 2 552 6 8 7 4 1
33144: 95 7 2 9 49 20
1128044: 8 9 2 1 9 84 752 2 2 40
2654256: 8 1 9 6 8 4 5 766 6 2 9 49
505: 55 1 9
2339768615183: 70 6 29 9 66 4 561 546
190334: 113 2 771 37
49303083: 2 4 453 3 7 8 8 2 7 76 6 3
203668202: 1 4 449 9 18 2 2 3 2 5 1 4
2362955: 8 7 4 7 254 87 7 8 5 28 2
732925: 2 3 5 31 2 5 9 407 9 4 3
392546: 3 925 2 9 5
4885587341: 58 43 995 9 836
188814: 3 8 143 55 54
51483: 62 4 34 8 6 110 55 681
578151515: 55 35 133 483 5 15
7135707998: 990 43 54 72 723 99 8
55235: 65 24 51 12 767
297596: 854 86 7 948 4
7263823: 2 4 7 8 5 83 215 823
6767327088: 7 28 8 4 3 1 751 409 9
404445698: 23 226 18 174 3 1 70
85601664: 64 527 6 9 47
69252114860: 89 8 165 79 488 2 700
61160878625: 8 82 1 635 8 5 733 25
22585: 718 7 4 997 899 3 582
32373857: 518 62 76 93 88 860
133226301085509: 9 10 61 9 5 878 563 8 3
5346191229: 2 4 165 7 3 56 1 4 287
4526312: 1 7 672 89 143 917
1266486: 4 44 7 11 26
56268324002: 963 89 900 70 849 2
20680: 5 89 11 4 5 1
5500: 2 53 5 4 5
48510660513: 3 1 7 9 735 2 2 55 3 510
38102883: 422 6 7 5 91 504 3 8 3
3406860: 57 399 9 236 54 45 2
276814230: 3 369 5 2 7 2 7 7 1 5 51
105973: 7 495 7 7 7 6 5 4 4 10 13
50302: 50 2 30 1 75
43099: 53 812 58 1 4
1977619113: 59 7 68 514 137
97237624: 52 1 729 788 158
1097770009: 6 8 6 6 82 9 7 372 2 3 7 7
12444030231: 85 5 4 8 2 6 3 86 7 4 2 29
1530828: 7 528 6 32 5 24 468
837216: 705 3 333 38 9
392372: 9 13 498 638 2
821801974: 6 8 6 7 11 267 752 74
1649: 5 3 8 8 706 314 498 3
14010888: 518 49 2 46 6 24
3443382417: 7 6 453 2 3 80 2 35 6 9
3285898: 79 54 9 5 77
1151380440: 887 926 84 7 6 30 7 3 2
462: 73 2 63 2 44
19445085: 8 8 3 165 723
1174239: 356 53 9 319
367893451235: 76 33 1 423 957 833 2
177989037434: 8 92 7 239 8 6 5 87 432
720: 4 54 662
5186429: 1 517 640 3 1
138658884: 358 888 483 652 3 41
65435761: 1 8 66 1 1 8 859 32 3 6 1
4367866058: 8 2 157 138 3 42 4 5 2 3
8486791015: 6 3 925 59 35 6 5 85 15
295200671: 3 1 7 6 6 492 2 3 3 3 4 2
21012907: 75 119 9 8 98 26 3 55 6
49703808: 1 574 328 4 66
7449353: 8 94 38 8 99 2
20412: 4 2 8 81 18
791398: 223 9 6 52 94 8 4 3 358
3806: 8 5 36 50 6
14414: 2 38 9 926 33 6 728
241160: 93 6 1 4 5 6 9 7 65 1 8 1
13578478: 193 7 3 65 479
5166605928: 1 225 71 83 771 75 3 4
559815: 55 91 7 1 6
91984: 389 3 9 995 8
54184984: 5 6 5 5 2 79 6 6 30 2 692
6973: 11 74 82
1329287962: 73 978 4 642 29
43396: 903 4 85 642 6
3204797: 45 666 256 106 843
86172580: 8 61 725 76 7
170875673280: 44 406 360 4 962
63001440: 9 7 71 3 4 1 4 9 141 2 3 6
20150: 40 630 9 8 3
1305168: 40 47 29 517 85 692
4954665: 8 7 220 83 88
86687651: 7 293 5 52 8 475 3 50
2390975268: 90 5 9 951 6 8 2 561 71
695550835: 8 134 8 2 39 68 5 7 85 1
172872006: 400 42 49 210 6
725727288: 707 9 9 727 16 13 1
174331920: 31 13 5 35 4 28
10969820: 8 4 5 5 9 70 6 431 2 7 3 8
2148818: 1 8 7 6 456 1 76 3 8 50 9
382: 73 3 64 93 6
65814188: 9 4 81 8 3 47 8 9
4400514499: 59 567 1 98 125 692 1
309181201108: 43 43 773 8 8 65 9 52
38743084: 6 80 7 82 789 4 387 2
51137437257: 6 1 98 1 3 4 2 4 94 91 5
511515: 70 1 65 491 22 883 8
3626095718733: 990 640 51 970 59
5453474712: 2 458 55 7 26 85
3311: 1 8 55 5 1 9 347 25 50
180727: 4 45 627 7 53 42
35561: 7 4 30 3 61
36478823: 7 8 42 2 4 1 3 572 3 9 95
49493959223: 899 89 55 92 23
56807: 81 7 99 2 4
5642010: 4 398 425 120 33
648340: 361 164 28 34 1 44
639: 3 18 77 6 51
16922240: 896 5 59 5 85 64
719650088: 77 895 50 185 4 8 2 2 2
4750137600: 989 612 5 2 480
159429600: 66 198 5 4 610
7366: 2 5 7 150 9 7
1014: 964 37 5 7 1 1
389753: 40 26 374 793
16908487612: 78 756 531 6 60 9 52
1357504896486: 451 893 6 1 4 751 8 9
183288569: 65 39 7 56 1 5 68
69731768: 8 8 4 76 7 4 4 3 8 32 14 4
59045412533: 275 8 6 6 446 82 885 8
18042: 50 358 92 50
2985889789: 90 829 35 2 4 97 91
162921208: 2 792 139 9 5 36 8 41
28856144771: 787 2 7 582 5 3 900 71
22198416: 991 1 4 56 13
153803335401: 760 7 81 2 4 667 5 40 1
817087: 3 773 88 4 799
1637731: 9 4 4 394 24 51 3 6 9 1 4
16205356271: 2 45 9 1 1 6 781 4 9 1 4 5
130703: 482 6 9 3 403
18951: 9 34 3 59 20 73 627
737164800: 3 472 72 85 9 8 600 3 9
132: 92 23 8 9
155953: 71 99 1 917 64
8936: 82 2 4 67 5
27207840386: 91 24 769 32 162
584272: 88 67 99 567 1
34335036557: 343 253 97 365 59
472890: 61 7 6 452 91
213651: 31 5 1 9 133 1 5 3 819 9
143583840: 438 7 625 3 6 6 9 2 6 4 5
459573224: 567 1 9 899 9 9 31 9 5 9
1401048: 6 1 4 6 957
1148265: 9 442 71 43 69 51
1869750036226: 8 34 7 89 82 4 6 3 62 28
508326573891: 637 2 6 95 8 7 6 3 75 91
120780: 5 439 68 3 4
19646: 7 32 573 7 652 3
8285208839: 911 34 8 876 836
10073987: 85 90 5 14 94 7
133435348: 9 1 770 3 57 12 1 40 9 3
173086560285: 21 61 42 2 17 21 3 288
4818816000: 94 80 5 9 40 356
237720: 37 5 45 2 28
34469: 1 389 57 46 70 26 3
21593286430: 27 8 264 776 32
40469459: 5 2 4 5 79 8 67 8 5 8 1 19
66970044: 89 641 11 9 83
2684254253: 599 60 534 9 50 85 5 3
75032493: 6 5 3 4 4 3 4 2 231 4 6 1
76544: 52 2 711 45
383054: 41 2 880 441 10 244
310245: 44 5 38 5 6 2 8 5 8 245
1354593: 2 1 15 43 95
2289: 9 3 5 10 6 7 5 2
8595816301: 854 15 543 16 301
423622372: 9 30 532 1 2 88 93 116
29916056996630: 357 5 3 475 7 173 836
325135169: 6 60 1 1 2 560 45 7 8 2 9
1131922: 6 90 3 9 9 435 479 2 6
144222850: 288 4 45 6 1 50
397440: 22 62 2 184 23 64
173755: 9 5 420 2 79 2
5995489: 9 57 2 680 4 5 394 8 89
35582840783: 9 526 809 56 83 5 3 2 3
12690529743: 6 1 2 3 576 8 14 3 9 6 4 6
11608677571: 415 6 31 7 26 399 2
4502: 696 5 4 79 1 2 75 2 860
17281456908: 532 4 358 9 6 358
6376294: 59 18 7 6 93
7257802: 5 8 192 2 3 6 1 8 2 3 30 9
1420: 8 96 2 502 2
23974424565: 5 7 946 6 1 5 4 2 5 431 7
19615864: 1 886 74 93 6 7 589 2
29560776951: 1 2 363 879 3 3 1 95 1
4004: 4 3 7 2 62 25 4
207200: 25 9 435 7 6 65 2
149806265416: 535 637 5 965 639 4 4
612209762: 9 976 96 726 98
2929914946215: 7 369 1 568 7 8 7 218
50494441: 340 4 68 2 7 39 359 1
8580181: 8 580 148 28 4
8255953367: 33 800 17 155 583
7113: 2 5 4 16 70 8 875 35 7 8
812634: 6 382 35 985 1 583 1
150880: 785 5 615 435 82
10437168: 12 9 21 95 46
22202: 4 8 8 8 29 5 67 74 7 66 2
2178: 6 72 4 2 4 9 7 3 5 3 117
162735: 5 3 542 711 444 1 95
3494721190: 8 32 5 8 4 6 174 6 7 3 8 8
115737939: 660 7 9 10 1 6 1 7 594 9
3087899: 73 5 705 6 2
1381145663: 62 275 472 700 9 81
37834537955: 301 4 5 13 881 356 2 2
850870: 347 2 46 53 8
99124639: 9 1 6 4 8 6 30 8 4 463 6 3
56103175: 15 55 2 85 6 4 9 1 682
12000: 6 9 43 78 5
2036295: 400 493 3 4 49 9 21 1 9
707458500: 62 71 696 8 15 925
150452: 941 39 68 846 4
43747408850: 7 6 62 35 3 4 5 4 884 1 8
6607137: 54 3 14 18 659
741438930: 848 8 5 97 9 30
94009: 81 145 8 49 1
14139025531: 2 7 8 3 585 58 4 3 3 5 6 7
21726199464: 2 64 95 2 7 304 41
406374: 5 2 6 2 832 636 8 3 8 80
1446847927: 27 641 382 81 7 8
1806349515: 16 91 2 677 89 6
49344129: 2 728 2 3 749 5 1 6 1 8 1
81411478: 49 49 94 175 3 416 7
80220096: 4 8 1 8 50 764 99
93915525: 5 13 2 8 64 77 7 9 3 55
648: 1 480 9 97 61
18609831: 36 894 37 2 6 235 8 5 9
3099916: 8 26 549 455 9
124865438: 7 81 55 7 1 2 286 7 119
6006: 59 3 6 5 11
1953: 9 217 1 1
287232: 75 9 268 4 12 17
383283459: 9 2 6 1 813 2 1 2 2 2 24 8
175700: 366 480 12 8 1
265600: 7 1 385 5 816 384 16
5037: 8 9 94 81 3
5873378: 5 1 2 4 587 485 1 2 3 1 1
6270: 9 774 1 8 6
3171240: 92 766 1 45
1617: 7 8 9 41 3 11 3
210338: 7 4 758 8 8 89 4 3 4 727
18818: 2 6 547 1 4 946 364 4
5215: 122 32 708 311 292
5227164702: 8 177 569 6 8 369
737836233964: 21 7 85 9 4 169 64 963
865769351: 7 7 8 9 987 548 47 5 7 1
7095746: 8 8 9 105 19 3 75 73
28380359: 421 90 1 749
4685514: 244 4 6 89 8
100128871: 36 41 8 99 2 1 13
88579640005: 161 709 8 163 1 1 97
1550: 754 794 2
173337535627: 345 90 8 2 8 3 783
17319937118: 90 1 818 6 25 94 10 3 7
21500652741: 2 1 4 9 93 4 8 64 88 7 42
1124648: 2 258 5 5 6 8 1 88 5 9 6 2
205978: 2 42 847 209 798
32683: 7 6 429 23 54
2355: 655 62 7 866 765
380801339: 50 8 238 2 1 1 7 4 4 1 6 1
79368: 1 62 64 2 7
2593163: 7 37 3 154 8
1100284: 665 1 9 447 3 94 53 6 4
4039457351: 738 9 9 3 8 2 2 91 6 4 4 8
1381331: 4 4 5 62 1 719 154 329
284982744: 1 1 1 3 4 9 326 96 549 5
34950300: 5 3 1 2 3 6 84 1 5 89 7 60
87134199: 7 51 61 2 4 5 6 2 4 65 98
657524: 66 2 105 7 92
193465994904: 8 621 489 4 84 2 7 936
707968469357: 2 7 9 288 133 8 97 3
17969322: 8 90 177 2 141
206424226: 58 5 52 569 3 3 63
247830: 9 66 676 55 6
7511: 4 33 1 7 29
117256: 12 7 98 2 54
82908106588: 2 7 40 49 9 528 4 8 1 2 7
9882713559: 7 7 911 2 626 82 5 560
26593920: 2 76 865 9 5 4 3 8 54 1
146552020: 6 154 4 6 94 5 3 40 2 7 6
522380: 39 4 558 6 92
7157823: 71 30 3 24 823
36096379862: 49 207 47 12 6 6 2 3 2
438163196: 5 6 3 8 1 6 50 3 82 6 49 2
75788091: 94 37 4 361 8 79 9
13740363: 49 8 3 883 91
102964310: 6 5 245 3 831 9 429 2 3
1167309: 3 8 4 3 4 97 688 3 342 9
245032998: 3 12 57 202 67 467
5792554662: 8 9 4 8 554 659
182448: 31 10 4 8 8 8 9 2 5 56 9 7
4258: 9 7 21 5 60
5536520: 285 161 273 14 4 55
117216: 16 2 407 8 2
162731016: 7 9 1 3 56 233 7 6 7 6 4 9
4654432831942: 5 96 8 4 4 9 8 4 9 3 194 1
358480856: 7 3 24 56 1 527 1 2 9 58
4436: 9 4 3 2 9 7 6 57 839 1 3 3
51362: 1 1 61 421
332015004: 81 979 45 6 9
128824843980: 734 6 449 499 777
158239447425: 9 651 8 1 8 2 9 3 7 424 1
152881: 7 41 531 7 7 470
7210508228153: 6 64 299 92 9 628 155
31095458: 446 4 92 34 91 77 1 7
95788548: 1 352 82 8 3 45 12 6 6
323651196: 702 8 86 60 6 946
2466540: 87 48 59 270 10
1172691210: 37 595 479 39 4 77
119904911: 1 249 960 6 5 846
270: 38 6 2 158 66
159388: 10 5 8 43 96 1
114256303: 9 494 3 7 7 9 407 7
171233596: 3 458 2 983 2 559 47 4
34261895: 1 5 59 6 72 18 93
1483532664585: 8 17 912 4 68 6 488 6 5
66541: 5 3 8 552 300 1
13722098: 3 87 937 48 209 8
4691700: 67 16 16 788 6 8 25
128522: 96 1 66 2 788 875 7
287100: 30 1 330 1 29
1123526: 162 64 35 9 12
4448: 887 5 9 2 2
152519760: 1 9 5 6 845 376 590 4 6
1466874: 1 396 2 2 5 38 718
403694417: 4 70 1 8 3 5 366 98 3 74
138216247: 6 774 443 52 9 4
631151360: 34 36 232 16 7 347
27524: 4 28 850 313 9
6901494: 92 75 73 76 4
9810: 2 4 5 7 9 5 7 1 1 12 13 6
4682675: 4 673 9 6 73
435336: 361 2 26 97 6
590206386: 205 8 14 209 947
73703: 915 8 1 499
732: 93 609 30
68109876: 902 9 43 73 978
29884235: 830 5 7 3 9 3 359
6650771531905: 62 2 7 556 4 2 6 24 667
2275227: 25 270 8 95 9
1505599204: 8 6 21 5 569 5 9 25 72 4
82083: 12 8 5 57 1 3
257732: 61 91 23 298 895 2
206338833295: 9 2 1 34 2 6 7 5 781 4 4 1
2324223623: 9 43 3 5 4 16 6 6 30 18 4
96520666: 6 74 291 7 747 42 7
442535: 2 55 6 33 4
156413: 494 315 707 58 34 4
531554874: 5 30 939 615 867 6
14849318: 341 97 34 2 997
18577885924: 6 1 6 1 4 553 3 44 5 925
21162: 824 921 2 37 6
4757240: 9 3 3 5 6 1 321 6 2 4 5
26040: 114 4 2 7 56
13068618012713: 6 57 3 947 4 5 4 1 2 712
5211: 2 63 8 8
1760833: 122 60 2 691 24 7
6484146: 403 252 621 746 63
2905675573: 36 32 8 75 576
101684: 7 95 39 1 3 3 6 1 8 2 9 9
250904300: 877 285 13 36 22
546002180: 2 78 350 188 3 294
3884: 1 24 31 5 9
96107256: 5 7 2 9 481 5 5 8 5 152 3
31645500: 51 730 850
443: 431 5 7
30382: 88 343 121 70 7
54445355: 11 3 885 15 88 537
43129917217: 214 6 393 73 6 46
1173400: 8 56 8 2 122 3 1 7
727170: 72 8 789 8 9
149404567: 5 8 1 3 9 514 765 5 6 3 8
11765952256325: 2 72 36 25 4 9 534 6 8 8
83461036: 670 692 29 2 2 9 600
1727166997: 5 51 820 236 35 6 991
1316886978: 5 231 2 558 4 9 2 7 5 5 3
471408: 169 1 681 8 31 1 69
6617: 96 73 67 28 6
60539365: 968 59 10 3 1 2 53 2 1 8
353263708: 7 24 5 6 2 3 6 8 295 3 7 4
6210993972: 18 874 7 94 6 7 27 2 9 4
36735: 31 15 79
298423944: 927 240 64 4 91 666
3115048840: 7 8 193 1 179 7 6 7 6 23
977: 786 43 1 8 141
582880536: 81 3 37 621 302
347952: 6 569 60 295 2
153441694949: 7 1 4 3 7 7 208 2 9 4 9 49
163694248684: 8 8 2 60 9 43 4 9 1 283
494236514: 34 4 2 46 57 77 9
39861: 4 7 79 7 4 9 4 964 417 9
3065480: 3 8 3 1 85 8
260700: 6 9 31 71 34 79
33701972: 3 2 8 88 7 5 6 5 493
1467984: 61 1 8 81 3 424 4 8 2 72
359134: 6 88 2 340 94
330: 1 6 59 5 4 1
4748: 5 4 22 62 34
204469035: 82 872 1 76 4 8 3 705
8478: 7 4 9 3 62 6 2 4 5 1 758 6
22591: 482 2 4 291 29
37950996: 4 54 93 644 7 4
434584: 584 140 3 6 4
9992187: 995 4 1 37 189
3207013821: 3 7 43 4 7 9 6 73 664 8 8
10080589059: 490 65 1 1 633 1 5 753
1013043: 2 59 70 7 335 9
86562: 2 36 896 499 59 9
48909: 9 89 3 6 830
2856181467: 6 905 526 71 75 7
1251744: 47 70 3 2 1 8 5 3 3 95 6 4
119692: 3 9 2 44 9
47253408: 20 8 6 32 32 69
458540: 3 679 672 161 75
947154221: 9 228 245 69 996 6 53
2538720206: 492 1 12 500 860 209
5012: 42 8 1 1 4
679209552: 460 1 90 341 18 904
612172: 39 96 477 164 6
16931: 568 2 9 16 529 1 8 1 3 8
91036608: 8 3 274 3 840 39 4 9 6 4
39683: 94 419 2 291 4
24476464: 5 3 48 4 7 3 71 70 8 17 4
27982662680: 8 1 18 59 89 20 479
1164: 8 14 1 42
5488047242: 6 860 8 47 240 3
20809336: 826 94 134 2 91 653
1413424: 13 22 76 8 65 62 2
76930: 4 260 59 70
5721: 7 50 23
35829197: 5 2 8 41 74 1 156 6 48
49619: 658 75 269
4033089817302: 147 7 3 222 774 273
7580088: 20 379 10 63 22
3600172922: 405 4 44 4 1 86 3 2 3 22
191964843181: 74 2 6 9 7 8 9 15 14 6 3 1
793215272: 9 6 97 818 15 656
64197140: 56 712 2 322 5
2896720806: 2 776 5 37 62 2 40 30 6
555886: 206 4 9 6 2 7 2 7 7 9 49 7
1237087: 1 732 169 2 8
7062363: 596 2 6 7 2 295 6 3 47
20735040: 5 715 4 29 8 5
61779067685: 4 39 60 46 48 66 1 5
702467: 777 3 490 249 38
471908563477: 594 3 352 6 115 5 79
708223: 88 8 4 22 3
536612544000: 46 885 619 768 30 25
28290: 11 63 380 3 2 157 8
3610965605: 878 33 28 58 90 55 8 8
64945132: 6 494 419 88 9 42 9
467864: 4 772 60 5 5 69 93 7 9 8
601020: 954 10 9 7
13757201: 42 935 35 350 451
107517: 633 424 18 1 5
606816441: 499 304 8 4 4 44
4081539268: 7 9 6 71 8 66 28 422
3840924780136: 57 241 74 6 671 138
2169263815300: 87 757 39 86 39 737 8
142140253: 202 4 690 2 53
63834: 787 9 3 3 86
580608: 49 7 2 644 768 1
1156213: 2 486 6 827 45 30
15309973132: 9 170 997 31 35
324351: 32 2 37 9 9 9 363
38490: 6 4 1 5 6
67409678: 96 29 87 7 16 8 423
1242: 78 4 454 4 1
41446548: 1 799 2 44 5 1 8 97 83 6
526892512645: 327 4 9 4 88 4 51 26 4 3
399019852: 9 283 9 3 71 4 828 4 7
59417893: 99 6 17 89 3
13016850: 75 7 165 35 77 7
254225563200: 8 98 60 5 817 3 441
74208: 72 13 681 7 96
480909: 4 9 8 5 1 3 1 78 29 14 8
183416: 312 7 8 1 6 31 39 454
1356: 5 5 98 81 315 2 835
2725632: 8 3 40 6 9 43 3 9 448
46877174: 58 40 1 9 8 19 58 172
5229: 6 549 694 10 569 664
62698125845: 2 8 72 1 59 37 182 843
34062: 813 29 6 541 84 6
39989600694298: 6 826 8 15 747 9 3 74 4
61809443: 1 642 6 5 4 8 5 5 1 4 638
60974222881: 66 8 88 41 8 1 6 2 8 30
31733680: 92 583 4 3 9 47 2 110
4243751006: 858 7 1 7 5 70 44 2 566
818321588: 9 3 8 8 6 3 619 5 4 7 61
441826: 770 9 9 7 4 7 9 6 1 4 4 20
2592497732: 7 4 1 82 2 3 7 9 4 737 82
112808: 1 5 1 236 8 882 5 7 4 4 2
73879: 738 36 42
160: 49 87 9 9 6
350312763: 405 8 3 8 8 9 5 9 7 4 2 7
35194: 37 307 71 61 662
34044006: 745 96 7 68 486
2182659: 859 6 7 417 6 918
212200: 6 66 59 7 9 4 98 9 862
5217957: 69 7 6 2 3 9 9 31 1 15 9
56946: 9 692 9 381 513
16404570432580: 346 878 2 7 7 6 8 9 58 3
3529315: 4 5 3 967 34 5 3 9 7 1 2
1611286: 1 824 7 88 5
1989: 47 2 8 4 11 17
289140827: 66 544 474 827 3
245987: 42 7 1 5 990
55544: 314 1 8 256 5 2 2 7 8 9
197: 8 8 6 7 43
48125625829: 3 45 11 1 5 576 49 826
1056184768: 27 2 37 39 9 556 6 82 4
7273400: 36 3 9 5 5 2 7 5 41 5 5 4
1184439754512: 76 39 5 7 53 8 19 51 2
659919992: 5 86 26 442 9 51 631
1488: 5 4 32 4 595 33 797 7
27442807: 761 9 36 990 7
332850201689: 76 4 795 93 994 9 548
1383172: 288 26 881 5 2
3128670: 9 701 6 9 763 5 2 1 8 6 5
67400940889: 991 17 3 2 35 2 23 4
152219721: 11 16 18 9 2 96 9
74970: 5 448 4 19 76 7 512 70
113424: 3 7 3 1 9 62 68 6
1635: 26 61 3 54 708
2922193: 8 91 1 9 364
662276171: 8 2 5 984 57 7 2 8
7523748364: 69 1 645 181 934 93 1
18896350006: 685 46 97 98 29 2
142121662: 8 294 6 90 10 160 6 4
570544: 5 705 45 1 1
771329: 55 1 172 48 326
16479828637: 5 6 51 8 6 1 9 99 57 6 4 5
7392: 40 9 2 4 11 1 4 2 40 1 1 3
46973: 56 1 5 9 671
6489357: 4 7 8 7 9 7 9 149 5 123
48848: 72 431 97 8 49
1622: 24 435 1 803 359
18174698: 677 9 65 242 4 9 7
110260316726: 737 3 27 149 3 16 726
44217591: 3 432 18 699 594
7421406399: 41 181 405 1 399
1641691441: 1 8 3 6 3 580 2 70 48 1 3
15169593474084: 2 27 362 8 10 834 1 84
13128338: 58 2 7 79 3 2 36 74
48109575: 31 9 3 1 3 418 8 97 389
78259125: 781 78 81 125
639039: 2 5 7 1 72 9 58 5 92 7 1 7
487956: 8 248 7 34 9 834 6 84
273739: 396 691 85 9 9
396470: 278 3 349 832 36 34
36577051: 914 7 55 76 721
18450: 9 7 9 2 22 92 765 2 9
24840211649: 9 2 41 2 3 84 4 4 3 60 86
350895218688: 8 1 8 7 512 943 7 927
10037001: 4 7 2 5 904 2 463 92 4 4
103811854: 802 40 40 404 8 2 1 9 3
4314699: 3 3 861 8 8 326 6 3 1 8 6
188: 65 6 2 3 43
4685347220: 9 6 304 478 93 72 21
2813616: 5 2 2 1 4 2 514 7 1 24 6
46964693: 10 8 1 98 847 594 9 83
20673316313078: 8 48 262 385 538 78
6788275: 2 70 46 310 38 184 92
1669684899: 16 900 49 372 99
95580: 2 3 54 3 90 6
21006: 5 415 60 8 671 7 1 1 6 3
1492494: 36 92 674 7 2 8 4 7 23 4
5285: 9 73 8 76 5
2115745: 19 89 424 33 4 411
21652829: 8 786 7 39 99 593 7 7 4
3185296595: 6 78 722 808 979 65
74810483: 3 9 5 8 34 6 511 3 52 31
4636: 3 42 579 2 4 5
2251608477: 25 2 6 290 4 84 6 4 4 7
519939: 846 9 7 434 81
14198101682: 84 9 6 8 683 90 4 62 15
4024180760: 75 4 1 8 816 4 8 2 8 4 6 4
1341759123: 9 2 485 33 23 354 3
13260: 4 8 9 310 1 7 593 1 13 1
24554932032: 236 9 54 1 4 4 320 2 7 6
6766: 4 5 957 7 4
685448379: 84 7 7 91 8 8 3 76
86993928: 2 9 840 93 930
1889442: 6 3 6 14 46 988 471 5 7
228050008: 61 178 6 3 6 50 7 70 18
7031559: 931 23 6 16 328 879
303077473: 232 72 378 6 8 5 7 7 78
2217338: 29 9 6 676 989
260719: 21 7 96 9 9
766868850: 845 35 672 63 2 3 1 45
405738630: 447 583 3 519 253
1462862520667: 3 26 69 2 2 9 5 7 59 665
14506615556: 266 3 908 64 928 772
48771863750: 9 59 92 75 409 7 750
15877524: 15 313 563 699 825
99: 23 68 5
6927: 3 100 159 430 7
2868775: 537 12 712 325 7
4441: 563 1 1 4 864 3 4 8 133
66275: 4 8 6 2 3 6 2 95 1 6 7 148
475793: 6 6 8 66 9
153942: 77 74 3 9 6 3
23122854: 61 10 3 320 36 52 1
2560805: 2 20 16 97 75 5
16828887: 117 93 307 54 8 6
80181: 1 534 1 82 8 3 31 6 7 3
287296497: 8 901 6 949 4 7 5 2
309561046: 6 2 30 1 2 5 3 2 69 2 18
478544: 598 11 8 6 48
76427538: 653 8 19 770 99 319
667296: 887 863 53 5 37
46887: 32 15 1 319 3 19 932
69166: 673 68 846 1 1 948 6
4577874: 87 370 7 730 71 73
63345725: 375 388 83 16 7 22
21871384: 51 325 52 350 4 2 73
819695819: 321 81 7 7 750 643 62
298545: 815 61 8 6 207
11723610: 21 6 1 35 64 52 3 5 6 10
1257164: 10 41 33 285 91
3409042: 8 56 3 38 86 7 5 2 89
32629024: 1 48 3 113 42 51 2 6
595289: 18 8 43 3 4 8 9 8 44 4 6 3
115839004217: 6 355 186 6 55 98 24
403829868080: 9 8 3 784 4 6 6 8 438 81
105: 6 2 5 7 14
69477835: 6 1 35 405 21 8 5 7 4 7 5
1144: 6 6 3 15 6 55 4
3801: 1 51 53 762 104 9 223
5521373235: 3 368 274 6 38 5 43 2
152324441069: 476 40 555 8 98 8 8 1
5651021807: 437 4 779 415 5
2840: 139 5 4 52 8
140994: 7 29 25 6 3 30 2 532 4
21954610664: 530 3 9 39 46 23 89 62
1670885: 6 274 54 5 885
1423556730: 1 87 972 2 5 8 2 5 3 10
14403093599: 580 94 714 37 2
10831103: 676 3 2 5 76 11 59 70 3
6211816367: 3 34 4 9 32 28 9 776 6
3584183: 84 54 9 79 30
29246400: 86 1 6 5 56 749 60 9 16
7848935528: 5 6 9 9 162 5 51 592 6 6
3549205: 51 906 54 9 65 85 6 85
60991691040: 55 7 63 2 7 2 8 9 90 97
38424: 168 35 432 92 1 3 2
6129304: 8 8 86 53 3 6 2 1 1 1 3 35
357635: 162 731 782 3 23
8608: 91 985 8
661739: 4 2 523 167 904 9
9473627: 6 174 1 4 2 5 4 9 5 90 45
673: 585 66 7 14 1
15129547092710: 22 849 42 3 32 9 9 709
261741896: 2 3 7 3 1 49 9 7 418 3 6 9
524610: 1 5 6 3 6 97 3 23 3 5 17 5
85278685: 8 527 86 8 8
1297: 4 54 249 229 758
6923467: 65 4 23 470
1113314: 5 755 3 8 6 9 3 3 9 3 536
12894111360: 63 2 942 351 95 6 1 8 8
79212000: 77 5 3 6 66 328 1 5 3 7 5
2758: 90 5 14 88 5
537644: 896 6 12 25 10
12625: 528 675 4 504 89
3708240060: 10 27 9 3 24 6 400 2 3 7
267446144: 793 50 50 1 3 842 8
244023: 7 6 27 3 3 20 942 5 37 8
177001872: 930 9 433 11 152 842
20416347: 9 9 3 4 8 5 3 139 6 3 8 1
133544160: 3 27 3 2 454 3 76 76 45
4811: 4 65 9 6 58 2 3
946450110057: 8 8 274 65 263 7 6 81
908336037: 9 801 9 7 90 5 6 1 5 4 1 2
17868897: 7 50 66 8 55 2 248
12996993: 926 3 6 54 9 1 730 2 73
47155063: 4 71 54 3 7 41 24
95223009: 7 165 443 8 45 3
20447786708: 7 66 7 4 777 7 9 702 6
19024: 51 74 5 2 152
934847: 822 98 9 5 847
4766263: 5 77 75 775 13
168767289: 840 11 1 1 1 36 198 9
3407322: 3 7 4 7 2 9 25 516 42 4 2
9530549257: 4 8 6 4 8 79 8 103 6 4 5 6
1730304: 59 76 5 40 1 35 4 4 2 8
207258242: 4 3 4 1 3 7 67 926 7 384
9321365115040: 6 214 1 9 5 3 410 5 6 8 5
5507765: 140 43 6 914 1
18137448: 25 70 9 47 27 812 1 15
39453587838: 268 391 7 54 3 49
10218744003: 75 694 40 27 50
1326269: 1 12 788 73 606 6 8 6 2
248566: 6 96 7 6 422 8
39237496416: 4 1 7 1 8 62 872 3 48 9 6
842285879: 842 285 697 6 176
281214058: 433 96 18 22 4 3 9 28
2699428486: 3 6 8 7 8 3 6 7 51 4 47 8
92281406: 26 6 766 81 405
50443200: 5 4 6 8 5 339 31
865688455: 3 317 7 6 95 6 9 7 5 3 18
673282920: 25 737 109 7 3 7 9 5 8
3855934: 55 982 5 22 6 604
102726: 38 917 72 2 6
1897272: 361 5 9 900 4 57 906 7
64516135358: 9 484 5 4 7 507 86 6 6
26238634: 225 86 678 2 34
371224: 31 187 64 62 154
238597: 9 9 343 7 851 46
23965014053: 47 93 625 1 757 8
19996130502: 1 1 8 6 9 213 6 669 7 5 3
11334686: 18 889 32 182 2 3
1062213: 531 2 213
4324: 83 2 974 601 421 2
247284576: 8 654 94 72 59 77
2312744: 15 9 96 274 4
235683002: 2 4 19 7 3 4 8 30 1 43
1008: 2 95 358 545 8
10235870: 2 2 9 5 3 7 9 9 4 73 1 195
2590: 940 2 31 13 666
67174779337: 91 45 31 39 460 529
3676536384: 448 675 94 2 206 59
54284: 1 1 1 7 353 8 718 5 7 7 4
986: 26 3 1 91 867
42685484: 39 4 9 39 8 51 7 1 1 17
39005: 8 28 13 877 29
34098587: 4 5 1 2 554 69 2 4 1 5 2
97902: 194 3 8 2 3 5 7 7
6758911: 29 1 1 9 463 3 7 4 5 9 7 8
13251598: 8 1 2 3 9 19 8 5 1 469 3
4050552: 974 4 89 3 652 51
372902: 5 8 3 2 8 44 2 8 9 80 81 3
35475443: 14 5 21 8 566 77 4 13 3
766: 9 7 7 88
365974: 6 11 39 552
11260857806: 31 41 7 93 6 68 85 5
11578: 1 36 32 5 8
2371: 8 5 10 74 611 58
740390456680: 739 499 891 45 6 68 1
8590798: 2 69 59 3 400
565: 45 491 29
73454669713: 734 54 571 98 711
9785129118: 35 7 17 24 3 137 1 6
8322: 912 8 9 35 7
20049271380: 886 1 9 5 6 399 14 4 45
1986906: 6 8 89 758 4 555 6
2668504: 61 927 9 3 30 3
10134: 302 5 33
88753182: 177 50 5 2 519 663
`
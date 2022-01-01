package day_05

const input = `
105,697 -> 287,697
705,62 -> 517,250
531,627 -> 531,730
21,268 -> 417,268
913,731 -> 271,89
214,697 -> 82,697
376,661 -> 376,177
519,859 -> 977,859
782,98 -> 184,98
612,179 -> 515,179
340,772 -> 352,784
111,863 -> 111,298
944,73 -> 594,73
465,21 -> 970,21
122,592 -> 111,592
975,975 -> 16,16
327,532 -> 561,532
811,618 -> 811,945
623,437 -> 623,202
380,591 -> 871,591
278,514 -> 125,667
797,946 -> 953,946
325,61 -> 484,61
450,422 -> 450,862
923,972 -> 119,972
813,141 -> 69,885
926,834 -> 926,687
137,564 -> 595,106
415,566 -> 274,566
726,354 -> 251,829
889,236 -> 470,236
282,376 -> 282,193
343,248 -> 932,248
790,918 -> 790,528
532,369 -> 222,369
15,378 -> 820,378
279,507 -> 279,719
641,68 -> 220,68
340,270 -> 340,680
939,364 -> 32,364
686,106 -> 568,106
919,365 -> 255,365
870,236 -> 879,227
322,397 -> 397,322
984,980 -> 350,980
392,864 -> 31,864
846,975 -> 243,372
253,981 -> 500,734
98,193 -> 280,11
477,460 -> 350,460
690,833 -> 48,191
469,409 -> 218,409
321,532 -> 321,106
868,341 -> 223,986
185,174 -> 801,790
256,658 -> 800,658
808,576 -> 931,576
959,913 -> 959,785
976,969 -> 47,40
891,931 -> 572,612
600,804 -> 866,804
149,368 -> 680,899
799,882 -> 157,882
803,214 -> 803,668
53,900 -> 940,13
424,800 -> 424,261
985,924 -> 80,19
158,194 -> 158,281
683,237 -> 683,341
493,482 -> 493,921
664,195 -> 664,824
689,405 -> 616,478
946,873 -> 846,873
977,988 -> 28,39
305,892 -> 662,892
891,27 -> 891,440
136,897 -> 35,897
948,458 -> 935,458
569,100 -> 599,100
542,292 -> 974,724
501,825 -> 104,428
875,872 -> 875,441
631,924 -> 43,336
874,846 -> 874,389
947,932 -> 81,66
75,480 -> 75,403
211,622 -> 211,482
344,904 -> 699,549
227,508 -> 698,508
677,774 -> 385,774
279,267 -> 391,155
294,801 -> 547,801
717,446 -> 614,549
490,903 -> 490,225
872,751 -> 278,751
580,163 -> 61,163
198,800 -> 389,800
147,728 -> 516,728
675,417 -> 675,752
147,544 -> 134,544
977,70 -> 164,883
349,976 -> 349,23
897,10 -> 14,893
602,349 -> 602,354
326,332 -> 355,332
53,331 -> 34,331
617,333 -> 466,333
661,537 -> 661,131
985,18 -> 20,983
953,580 -> 953,124
70,363 -> 74,363
448,38 -> 141,38
957,175 -> 957,634
88,316 -> 88,899
231,94 -> 857,720
643,566 -> 643,832
724,955 -> 243,474
368,521 -> 537,521
649,245 -> 406,245
92,304 -> 399,304
978,491 -> 819,491
99,637 -> 765,637
243,159 -> 803,719
139,756 -> 305,756
815,226 -> 79,962
317,562 -> 491,562
783,95 -> 783,277
207,321 -> 133,321
752,136 -> 185,703
752,990 -> 752,433
282,841 -> 466,841
314,31 -> 314,829
637,873 -> 637,854
60,746 -> 563,243
646,566 -> 119,39
260,475 -> 124,339
603,647 -> 327,647
990,202 -> 342,202
981,620 -> 606,620
475,352 -> 313,352
184,497 -> 143,497
130,929 -> 329,929
779,111 -> 779,975
892,960 -> 11,79
37,984 -> 919,102
589,794 -> 589,548
665,668 -> 385,668
668,301 -> 281,301
860,122 -> 623,122
18,914 -> 782,150
691,150 -> 25,150
117,439 -> 462,439
926,695 -> 926,651
907,644 -> 708,644
545,120 -> 229,120
181,659 -> 181,820
362,543 -> 575,330
603,531 -> 603,142
754,404 -> 754,678
703,551 -> 450,551
794,137 -> 581,137
866,288 -> 327,827
676,613 -> 676,470
874,130 -> 23,981
132,288 -> 360,288
706,147 -> 706,433
734,646 -> 588,500
641,386 -> 598,343
743,726 -> 79,62
308,192 -> 859,192
858,125 -> 603,125
694,199 -> 653,240
251,407 -> 79,407
254,337 -> 254,310
586,850 -> 17,281
937,989 -> 17,69
503,784 -> 584,784
17,97 -> 906,986
909,987 -> 23,101
11,465 -> 953,465
645,862 -> 251,862
741,488 -> 856,488
488,123 -> 488,641
720,775 -> 79,775
228,105 -> 702,105
344,804 -> 873,275
953,848 -> 669,564
188,76 -> 524,76
473,852 -> 137,852
515,14 -> 515,183
362,654 -> 362,335
76,73 -> 969,966
987,743 -> 468,743
912,28 -> 912,31
464,247 -> 380,331
171,20 -> 171,863
855,653 -> 855,941
505,415 -> 505,808
947,543 -> 947,821
907,365 -> 726,365
475,563 -> 475,63
927,679 -> 773,679
938,77 -> 26,989
345,909 -> 299,909
46,22 -> 972,948
197,735 -> 288,735
552,748 -> 756,952
946,180 -> 946,695
956,779 -> 216,779
120,105 -> 950,935
924,902 -> 35,13
530,49 -> 451,128
491,693 -> 340,693
533,774 -> 623,864
177,618 -> 177,123
543,114 -> 637,114
503,585 -> 344,585
34,836 -> 34,625
618,802 -> 212,396
863,678 -> 349,678
26,850 -> 768,108
99,67 -> 988,956
11,902 -> 871,42
658,749 -> 507,900
967,178 -> 218,927
671,247 -> 671,525
421,985 -> 541,865
279,639 -> 754,164
627,747 -> 627,290
77,66 -> 977,966
177,282 -> 617,722
400,444 -> 451,393
540,152 -> 540,888
521,196 -> 36,196
32,590 -> 32,537
145,613 -> 279,747
45,428 -> 45,12
785,956 -> 785,728
205,507 -> 205,539
117,12 -> 117,221
395,17 -> 479,17
104,881 -> 933,52
918,716 -> 570,716
121,621 -> 937,621
516,773 -> 516,917
311,605 -> 311,168
611,185 -> 611,976
373,80 -> 373,295
987,295 -> 515,295
416,717 -> 416,121
251,508 -> 196,453
498,824 -> 428,754
956,818 -> 153,15
266,272 -> 266,748
769,312 -> 769,387
604,766 -> 184,766
656,934 -> 520,934
224,771 -> 162,771
588,395 -> 133,395
219,489 -> 219,948
67,42 -> 979,954
684,109 -> 920,345
168,895 -> 762,301
761,953 -> 59,953
583,408 -> 592,399
129,48 -> 931,48
694,76 -> 404,76
808,380 -> 808,886
643,165 -> 643,757
714,543 -> 714,913
258,550 -> 295,550
400,857 -> 400,38
267,573 -> 267,779
124,182 -> 255,51
399,981 -> 552,981
197,803 -> 197,275
791,706 -> 791,373
500,664 -> 924,664
177,171 -> 177,935
703,43 -> 696,43
265,849 -> 889,225
847,324 -> 661,324
369,965 -> 369,780
169,965 -> 935,199
742,540 -> 742,355
210,854 -> 204,854
58,281 -> 954,281
858,793 -> 666,793
276,156 -> 733,613
537,538 -> 80,81
985,10 -> 14,981
79,31 -> 692,644
77,41 -> 77,502
684,150 -> 17,817
295,785 -> 920,785
171,579 -> 171,16
763,754 -> 763,86
719,573 -> 719,71
183,708 -> 227,708
826,952 -> 835,952
124,914 -> 975,63
807,704 -> 653,704
140,468 -> 140,874
408,330 -> 408,291
501,958 -> 501,302
834,505 -> 686,357
267,76 -> 267,526
18,88 -> 863,933
147,188 -> 147,454
922,733 -> 277,733
509,259 -> 957,259
614,765 -> 238,765
77,54 -> 77,252
591,532 -> 591,384
539,574 -> 729,384
347,158 -> 347,10
389,988 -> 989,988
696,571 -> 662,605
656,207 -> 656,883
802,446 -> 802,693
121,35 -> 121,66
967,738 -> 949,738
12,86 -> 809,883
96,167 -> 758,829
790,42 -> 790,549
14,987 -> 986,15
363,689 -> 363,386
148,148 -> 807,807
891,899 -> 891,710
445,678 -> 445,464
649,426 -> 649,452
641,378 -> 967,378
580,220 -> 300,220
376,789 -> 376,572
770,551 -> 647,428
651,692 -> 399,692
432,385 -> 432,835
242,48 -> 512,48
955,612 -> 955,520
926,568 -> 938,556
626,836 -> 626,266
973,982 -> 39,48
64,32 -> 64,653
503,444 -> 641,444
593,306 -> 11,888
287,138 -> 287,891
529,886 -> 529,826
217,320 -> 217,875
11,988 -> 989,10
291,30 -> 488,30
864,945 -> 113,194
550,501 -> 550,89
269,474 -> 269,40
953,394 -> 908,394
451,983 -> 451,293
135,121 -> 455,121
30,35 -> 915,920
31,451 -> 31,936
300,715 -> 42,973
577,459 -> 577,700
291,539 -> 456,539
373,449 -> 855,449
222,136 -> 358,136
206,14 -> 206,211
977,577 -> 977,535
183,723 -> 183,900
888,905 -> 821,905
51,301 -> 388,301
859,594 -> 859,227
767,343 -> 767,472
36,897 -> 565,897
450,481 -> 855,481
137,401 -> 137,643
771,276 -> 771,61
767,144 -> 767,562
212,111 -> 978,877
841,117 -> 234,724
975,104 -> 263,104
839,408 -> 839,588
122,50 -> 911,839
748,208 -> 748,929
230,305 -> 645,305
107,324 -> 175,256
726,339 -> 726,968
780,127 -> 664,11
392,148 -> 392,133
228,607 -> 228,689
469,379 -> 739,379
797,851 -> 841,895
896,494 -> 896,568
351,950 -> 566,950
593,387 -> 492,488
939,664 -> 843,664
463,159 -> 197,159
164,265 -> 164,16
164,147 -> 510,493
989,988 -> 11,10
98,676 -> 693,676
118,384 -> 118,544
220,502 -> 220,593
530,437 -> 802,437
321,29 -> 321,819
438,118 -> 438,531
268,128 -> 802,128
602,770 -> 602,183
841,58 -> 846,63
582,371 -> 592,361
174,163 -> 296,163
927,268 -> 927,391
579,280 -> 12,847
52,951 -> 52,772
645,203 -> 985,203
725,119 -> 725,367
155,112 -> 779,736
988,44 -> 320,712
438,463 -> 914,463
193,948 -> 292,948
217,398 -> 638,398
70,553 -> 465,158
271,262 -> 867,262
964,576 -> 442,54
253,67 -> 972,67
537,507 -> 290,260
537,645 -> 213,321
366,130 -> 913,677
834,283 -> 834,523
858,825 -> 858,391
146,60 -> 146,701
865,909 -> 162,206
503,628 -> 326,628
49,101 -> 583,101
692,17 -> 692,218
704,744 -> 210,744
144,434 -> 587,434
630,393 -> 630,870
606,616 -> 606,330
41,83 -> 916,958
80,341 -> 706,967
426,683 -> 426,173
919,962 -> 499,962
442,49 -> 442,970
740,378 -> 498,378
563,196 -> 563,442
222,76 -> 614,76
398,451 -> 851,451
62,50 -> 243,50
775,114 -> 775,234
650,901 -> 650,195
164,10 -> 164,149
127,751 -> 67,751
122,674 -> 780,674
325,652 -> 70,652
944,908 -> 99,63
40,985 -> 977,48
946,21 -> 126,841
872,906 -> 872,136
365,288 -> 827,750
348,935 -> 244,935
371,963 -> 499,963
816,595 -> 392,171
953,673 -> 953,585
223,612 -> 223,362
327,423 -> 553,649
661,693 -> 258,693
10,838 -> 10,859
985,814 -> 985,25
331,529 -> 87,529
611,460 -> 355,460
928,426 -> 748,426
540,172 -> 365,347
57,45 -> 57,129
20,861 -> 628,253
460,474 -> 297,311
549,876 -> 131,876
748,197 -> 287,658
639,137 -> 741,137
917,35 -> 917,273
482,333 -> 975,826
176,817 -> 89,730
894,418 -> 806,418
555,227 -> 349,433
317,33 -> 432,148
93,988 -> 93,479
635,300 -> 870,300
301,437 -> 301,760
660,548 -> 660,909
696,18 -> 60,18
231,787 -> 165,787
500,242 -> 371,242
88,126 -> 405,126
983,941 -> 61,19
242,519 -> 242,489
519,957 -> 926,550
606,181 -> 606,432
873,216 -> 851,194
880,924 -> 880,844
321,119 -> 801,599
963,392 -> 726,155
190,655 -> 190,305
542,676 -> 542,819
`

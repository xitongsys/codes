all:


clean:
	rm -rvf a b c d e f g h i j k l m n o p q r s t u v w x y z
	cp -rvf sample a
	cp -rvf sample b
	cp -rvf sample c
	cp -rvf sample d
	cp -rvf sample e
	cp -rvf sample f
	cp -rvf sample g
	cp -rvf sample h
	cp -rvf sample i
	cp -rvf sample j
	cp -rvf sample k
	cp -rvf sample l
	cp -rvf sample m
	cp -rvf sample n
	cp -rvf sample o
	cp -rvf sample p
	cp -rvf sample q
	cp -rvf sample r
	cp -rvf sample s
	cp -rvf sample t
	cp -rvf sample u
	cp -rvf sample v
	cp -rvf sample w
	cp -rvf sample x
	cp -rvf sample y
	cp -rvf sample z

.PHONY: a b c d e f g h i j k l m n o p q r s t u v w x y z

*:
	g++ ./$@/sol.cc -o ./$@/sol.exe
	./$@/sol.exe < ./$@/in.txt

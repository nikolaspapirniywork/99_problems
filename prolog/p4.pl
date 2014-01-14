 %%(*) Find the number of elements of a list.

my_length(0, []).
my_length(R, [_|Xs]) :-
	my_length(R1, Xs),
	R is 1 + R1.

:- begin_tests(p4).
:- use_module(library(lists)).
	test(length_1) :-
		my_length(5, [1,2,3,4,5]).
	test(length_2) :-
		my_length(1, [1]).
	test(length_3) :-
		my_length(0, []).
:- end_tests(p4).

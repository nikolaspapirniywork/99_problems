%% (*) Find the last element of a list.
%% Example:
%% ?- my_last(X,[a,b,c,d]).
%% X = d

my_last([], []).
my_last(X, [X]).
my_last(R, [_|Xs]) :-
	my_last(R, Xs).

:- begin_tests(p1).
:- use_module(library(lists)).
	test(my_last_1) :-
		my_last(d, [a,b,c,d]).
	test(my_last_2) :-
		my_last(d, [d]).
	test(my_last_3) :-
		my_last([], []).

:- end_tests(p1).
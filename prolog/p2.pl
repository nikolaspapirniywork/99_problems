%% Find the last but one element of a list.
last_but_one(_, []) :- fail.
last_but_one(_, [_]) :- fail.
last_but_one(R, [_|Xs]) :-
	last_but_one(R, Xs).
last_but_one(X, [X|[_]]).

:- begin_tests(p2).
:- use_module(library(lists)).
	test(last_but_one_1) :-
		last_but_one(c, [a,b,c,d]).
	test(last_but_one_2, [fail]) :-
		last_but_one(_, [d]).
	test(last_but_one_3, [fail]) :-
		last_but_one(_, []).
	test(last_but_one_4) :-
		last_but_one(a, [a,b]).
:- end_tests(p2).
%% Find the K'th element of a list.
%% The first element in the list is number 1.
%% Example:
%% ?- element_at(X,[a,b,c,d,e],3).
%% X = c

element_at(_, _, El) :-
	El < 0,
	fail.
element_at(X, [X|_], 1).
element_at(R, [_|Xs], El) :-
	Next_el is El - 1,
	element_at(R, Xs, Next_el).


:- begin_tests(p3).
:- use_module(library(lists)).
	test(element_at_1) :-
		element_at(3, [1,2,3], 3).
	test(element_at_2, [fail]) :-
		element_at(_, [1,2,3], 4).
	test(element_at_3, [fail]) :-
		element_at(_, [1,2,3], -3).
:- end_tests(p3).
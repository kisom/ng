Functional Spec for ng
======================

Summary
-------

ng is a command line tool for generating a ninja_ config file.

.. _ninja: https://ninja-build.org/

Ninja is fast and this provides a tool for declaratively describing
how to build the project. It's a quick project that gives me exactly
what I want. There are some simple Python tools, but again --- this
is a quick project that gives me exactly what I want.


Usage sketches
--------------

Given a YAML file named "build.yaml" like::

    compilers:
      cxx: clang++
      cc: clang

    targets:
      cc:
	ch01ex01: []
	ch01ex02:
	  - ch01ex04.cc
      c:
	vm:
	  - vm.c
	  - stack.c
	  - isa.c

A user should be able to run ``ng``, which will find ``build.yaml``
and produce ``build.ninja`` automatically.


Requirements & Assumptions
--------------------------

This assumes a straightforward mapping of source files to build
targets, and is aimed at C/C++.

System Design
-------------

TODO: fill this in
^^^^^^^^^^^^^^^^^^

This is the "meat" of the spec.

1. What are the major components? What do they do?
2. How do they work together? You'll probably want at least one
   diagram in here.
3. What other components/systems do they depend on?

NOTE: You may or may not choose to go into the object model,
specific fields and datatypes, (see the module documentation
template). If you think it would be useful to get feedback on those
things, by all means include them. Remember, it needn't be set in
stone.

Supportability
--------------

TODO: fill this in
^^^^^^^^^^^^^^^^^^

1. What are expected failure scenarios? When they happen:

   + If it's a service and it goes down, how will you know? How will
     you get it started back up?
   + If it's a program that has data, how can the data be corrupted?
     How will corrupt data be recovered from?

2. What platforms will it run on? What sort of cross-platform support
   is needed?
3. What are the packaging requirements?

Security
--------

TODO: fill this in
^^^^^^^^^^^^^^^^^^

What are the major security concerns around the project? How will they
be addressed? If this is security software, what is the threat model?

Project Dependencies
--------------------

TODO: fill this in
^^^^^^^^^^^^^^^^^^

What does this depend on? Justify those dependencies.

Open Issues
-----------

TODO: fill this in
^^^^^^^^^^^^^^^^^^

It's rare that a spec fully buttons-up every conceivable part of the
functionality before implementation. What are the "known unknowns"?

Milestones
----------

TODO: fill this in
^^^^^^^^^^^^^^^^^^

What are some reasonable milestones for the project? Try for
milestones that produce working solutions that lack functionality. A
bare bones MVP is a fantastic milestone.  Rule of thumb: if you
can't make an 80% confident estimate for the date of your very first
milestone, this spec is probably not complete.

Review History
--------------

TODO: fill this in
^^^^^^^^^^^^^^^^^^

This may not be applicable, but it's usually nice to have someone else
sanity check this.

Keep a table of who reviewed the doc and when, for in-person
reviews. This document should have at least 1 in-person
review.

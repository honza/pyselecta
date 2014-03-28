pyselecta
=========

A fuzzy text selector for files and anything else you need to select

Inspired by Gary Bernhardt's `selecta`_.

.. _selecta: https://github.com/garybernhardt/selecta

Unlike selecta, pyselecta doesn't include any interactive functionality.  You
give it a list of options and a search term and it returns matches.

Installation
------------

::

    $ pip install pyselecta

This will give you the ``pyselecta`` shell command.

Usage
-----

::

    $ find . -type f | pyselecta "models"

This will return a newline separated list of file paths with the best candidate
on top.

::

    $ find . -type f | pyselecta


Omitting the search term just copies the stdin to stdout.

Vim integration (ctrlp)
-----------------------

Unlike the original selecta, pyselecta works in MacVim, too.

Add this to your ``vimrc``.

::

    let g:selecta_path = "pyselecta"

    function! SelectaMatch(items, str, limit, mmode, ispath, crfile, regex)
        let cachefile = ctrlp#utils#cachedir().'/selecta.cache'

        if !( filereadable(cachefile) && a:items == readfile(cachefile) )
            call writefile(a:items, cachefile)
        endif

        if !filereadable(cachefile)
            return []
        endif

        let cmd = "cat ".cachefile." | ".g:selecta_path." ".a:str
        return split(system(cmd), "\n")

    endfunction

    let g:ctrlp_match_func = {'match' : 'SelectaMatch' }


License
-------

BSD, short and sweet

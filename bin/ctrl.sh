#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"


function start()
{
    ${DIR}/livingserver &
}

function stop()
{
    killall -9 livingserver
}



case C"$1" in
    Cstart)
        start
        ;;
    Cstop)
        stop
        ;;
    Crestart)
        stop
        start
        ;;
    C*)
        echo "Usage: $0 {start|stop|restart}"
        ;;
esac

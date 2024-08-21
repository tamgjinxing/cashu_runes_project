#!/bin/bash

bitcoin_cookie_file_path="/Users/a001/bitcoin/data/testnet3/.cookie"

command=$1
target_address=$2
runes_name=$3
runes_count=$4
fee_rate=$5

wallet_name=$2

cd /Users/a001/bitcoin/ord-0.18.5

command_pre="./ord --cookie-file $bitcoin_cookie_file_path --testnet wallet "

main_command=""

send(){
    if [ -z "$target_address" ]; then
        echo "Error: The receive address is empty."
        return
    fi
    if [ -z "$runes_name" ]; then
        echo "Error: The runes name is empty."
        return
    fi
    if [ -z "$runes_count" ]; then
        echo "Error: The runes count is empty."
        return
    fi

    if [ -z "$fee_rate" ]; then
        echo "No fee rate provide.  use default fee rate: 30"
        fee_rate = 30
    fi
    
    main_command="$command_pre $command --fee-rate $fee_rate $target_address $runes_count:$runes_name"
    run
}

balance(){
    main_command="$command_pre --name $wallet_name $command"
    run
}

transactions(){
    main_command="$command_pre --name $wallet_name $command"
    run
}

receive(){
    main_command="$command_pre --name $wallet_name $command"
    run
}

create(){
    main_command="$command_pre --name $wallet_name $command"
    run
}

run(){
    echo $main_command
    # eval $main_command
}


case "$1" in

    send)
        send
        ;;
    
    receive)
        receive
        ;;

	balance)
        balance
		;;

	transactions)
        transactions
		;;
    
    create)
        create
        ;;
esac
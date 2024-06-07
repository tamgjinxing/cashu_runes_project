#!/bin/bash

bitcoin_cookie_file_path="/Users/tangjinxing/work/bitcoin/.cookie"
bitcoin_data_dir="/Users/tangjinxing/work/bitcoin/blocks"
server_url="http://localhost:8088"

command=$1
target_address=$2
runes_name=$3
runes_count=$4

wallet_name=$2

cd /Users/tangjinxing/work/bitcoin/ord-0.18.5

command_pre="./ord --bitcoin-data-dir $bitcoin_data_dir --cookie-file $bitcoin_cookie_file_path wallet --server-url $server_url"

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
    fee_rate=10
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
    # echo $main_command
    eval $main_command
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
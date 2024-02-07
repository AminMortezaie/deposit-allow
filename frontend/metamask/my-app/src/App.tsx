import React from 'react';
import './index.css';
import {useMetaMask} from "metamask-react";
import {Web3} from "web3";
import Header from "./components/Header";

//@ts-ignore
const provider = window.ethereum

function App() {
    const {status, connect, account, chainId, ethereum} = useMetaMask();

    const web3 = new Web3(provider);
    // console.log(web3)
    return (
        <header className="bg-blue-500 text-white ">
            <Header/>
            {/*<div className="container mx-auto flex justify-between items-center">*/}
            {/*    <div>*/}
            {/*        <h1 className="text-xl font-semibold">Deposit Allow</h1>*/}
            {/*    </div>*/}
            {/*    <div>*/}
            {/*        {status === "connected" && (*/}
            {/*            <div className="flex items-center">*/}
            {/*                <span className="font-semibold">{account}</span>*/}
            {/*            </div>*/}
            {/*        )}*/}
            {/*        {status === "notConnected" && (*/}
            {/*            <button*/}
            {/*                onClick={connect}*/}
            {/*                className="bg-white text-blue-500 hover:text-blue-700 font-bold py-2 px-4 rounded mr-2"*/}
            {/*            >*/}
            {/*                Connect to MetaMask*/}
            {/*            </button>*/}
            {/*        )}*/}
            {/*        {status === "initializing" && (*/}
            {/*            <div className="text-gray-200">Initializing...</div>*/}
            {/*        )}*/}
            {/*        {status === "unavailable" && (*/}
            {/*            <div className="text-red-300">MetaMask not available :(</div>*/}
            {/*        )}*/}
            {/*        {status === "connecting" && (*/}
            {/*            <div className="text-gray-200">Connecting...</div>*/}
            {/*        )}*/}
            {/*    </div>*/}
            {/*</div>*/}
        </header>)
}

export default App;

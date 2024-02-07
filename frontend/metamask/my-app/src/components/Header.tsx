import React, {useEffect, useState} from 'react';
import {SiweMessage} from "siwe";
import {BrowserProvider} from "ethers";
import {Web3} from "web3";

const BACKEND_ADDR = "http://localhost:3000";
const domain = window.location.host;
const origin = window.location.origin;
//@ts-ignore
const provider = new BrowserProvider(window.ethereum);
//@ts-ignore
const web3 = new Web3(window.ethereum);

const Header = () => {
    const [authenticated, setAuthenticated] = useState(false);
    const [address, setAddress] = useState('');

    useEffect(() => {

        const fetchData = async () => {
            try {
                // Make an API call to your backend
                const response = await fetch(`${BACKEND_ADDR}/login`, {mode: 'no-cors'}); // replace with your actual API endpoint

                if (response.ok) {
                    const data = await response.json();

                    // Assuming your API response contains address and authenticated state
                    setAddress(data.address);
                    setAuthenticated(data.authenticated);
                } else {
                    // Handle errors if needed
                    console.error('Error fetching data:', response.statusText);
                }
            } catch (error) {
                console.error('Error fetching data:', error);
                // Handle errors if needed
            }
        };

        fetchData();
    }, []);

    const handleSignInClick = async () => {
        try {
            const signer = await provider.getSigner();
            console.log("im going on1")
            const address = await signer.getAddress();
            console.log("im going on2", address)

            const message: string = address;

            console.log("message::", message);
            const signature = await signer.signMessage(message);
            console.log("sig::", signature);

            const verifyRes = await fetch(`${BACKEND_ADDR}/verify`, {
                method: 'POST',
                mode:'no-cors',
                headers: {
                    'Content-Type': 'application/json',
                },

                body: JSON.stringify({ message, signature }),
                credentials: 'include',
            });
            const verificationResponse = await verifyRes.text();
            // Do something with the verification response
            console.log(verificationResponse);
        } catch (error) {
            console.error('Error signing in with Ethereum:', error);
        }
    };

    return (
        <header className="bg-gray-800 text-white p-4 fixed top-0 w-full">
            <div className="container mx-auto flex justify-between items-center">
                <div>
                    <p className="text-xl font-semibold">Your Company Logo/Name</p>
                </div>
                <div>
                    <p className="text-right">{address}</p>
                    {!authenticated && (
                        <button className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
                                onClick={handleSignInClick}>
                            Sign In
                        </button>
                    )}
                </div>
            </div>
        </header>
    );
};

export default Header;

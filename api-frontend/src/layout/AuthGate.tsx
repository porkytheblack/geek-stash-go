/* eslint-disable react-hooks/exhaustive-deps */
import { useUser } from '@auth0/nextjs-auth0/client'
import React, { useEffect } from 'react'
import { pageProps } from '@/global/types'

interface IProps {
    onAuthenticated: (layout: 'home' | 'dashboard') => void,
    onNotAuthenticated: () => void,
}

function AuthGate(props: IProps & pageProps) {

    const { onAuthenticated, onNotAuthenticated, authOnly, layout } = props
    const { checkSession, isLoading, error, user } = useUser()

    useEffect(()=>{
        if(authOnly){
            if(isLoading) return ()=>{}
            if(error)  return onNotAuthenticated()
            if(user){
                onAuthenticated(layout || 'main')
            }else{
                onNotAuthenticated()
            }
        }else{
            onAuthenticated('home')
        }
        
    }, [isLoading, error, user])

  return (
    <div className="grid w-screen h-screen items-center justify-center">
        <p className="text-2xl text-blue-700" >
            { isLoading ? "Loading..." : error ? "An Error Occured" : user ? "Logging you in a sec...": "Redirecting..."}
        </p>
    </div>
  )
}

export default AuthGate
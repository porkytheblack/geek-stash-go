import { pageProps } from '@/global/types'
import { useRouter } from 'next/router'
import React, { useState } from 'react'
import AuthGate from './AuthGate'
import Dashboard from './dashboard'
import Home from './home'

type IProps = {
    children: React.ReactNode,
    pageProps: pageProps
}

interface IReducer {
    layout: 'home' | 'dashboard',
    processed: boolean
}

const initialState: IReducer = {
    layout: 'home',
    processed: false
}

const reducer = (state: IReducer, action: {type: string, payload: any}) => {
    switch (action.type) {
        case 'SET_LAYOUT':
            return {
                ...state,
                layout: action.payload
            }
        case 'SET_PROCESSED':
            return {
                ...state,
                processed: action.payload.processed,
                layout: action.payload.layout
            }
        default:
            return state
    }
}

function Layout(props: IProps) {
    const {push} = useRouter()
    const [{
        layout,
        processed
    }, dispatchAction] = React.useReducer(reducer, initialState)
    const { children, pageProps } = props

  return (
    <div className="grid grid-rows-1 w-screen min-h-screen h-full bg-black ">
        {
            !processed ? <AuthGate
                authOnly={pageProps.authOnly}
                layout={pageProps.layout}
                onAuthenticated={()=>{
                    dispatchAction({
                        type: 'SET_PROCESSED',
                        payload: {
                            processed: true,
                            layout: pageProps.layout
                        }
                    })
                }}
                onNotAuthenticated={()=>{
                    dispatchAction({
                        type: 'SET_PROCESSED',
                        payload: {
                            processed: true,
                            layout: 'home'
                        }
                    })
                    push("/")
                }}
            />
            : (
                layout === 'home' ? (
                    <Home>
                        {children}
                    </Home>
                ) : (
                    <Dashboard>
                        {children}
                    </Dashboard>
                )
            ) 
        }
    </div>
  )
}

export default Layout
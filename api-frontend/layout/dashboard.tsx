/* eslint-disable react-hooks/exhaustive-deps */
import Button from '@/components/Buttons/Button'
import LinkButton from '@/components/Buttons/LinkButton'
import Link from 'next/link'
import { useRouter } from 'next/router'
import React, { useEffect, useState } from 'react'

type IProps = {
    children: React.ReactNode
}



function Dashboard(props: IProps) {
    const { push, events } = useRouter()
    const { children } = props 

    const [pathChangeLoad, setPathChangeLoad] = useState(false)

    useEffect(()=>{
        events.on('routeChangeStart', (url)=>{
            if(["/dashboard", "/dashboard/keys", "/dashboard/usage"].includes(url)){
                setPathChangeLoad(true)
            }else{
                setPathChangeLoad(false)
            }
        })
        events.on('routeChangeComplete', (url)=>{
            setPathChangeLoad(false)
        })
    }, [])
    
  return (
    <div className="grid grid-cols-5 w-full  h-full divide-red-700 ">
        <div className="col-span-1 h-full bg-[#101016] py-2 px-2 ">
            <div className="flex p-5 flex-col items-center">
                <p className="text-xs text-green-700 text-center font-semibold">
                    geek-stash-go
                </p>
            </div>
            <div className='grid h-full justify-between grid-cols-1' >
                <div className="grid col-span-1  gap-1 grid-rows-sidebar ">
                    <Button
                        onClick={()=>{
                            push("/dashboard")
                        }}
                        buttonType="primary"
                    >
                            Dashboard
                    </Button>
                    <Button
                        onClick={()=>{
                            push("/dashboard/keys")
                        }}
                        buttonType="primary"
                    >
                            Api Keys
                    </Button>
                    <Button
                        onClick={()=>{
                            push("/dashboard/usage")
                        }}
                        buttonType="primary"
                    >
                            Api Usage
                    </Button>
                </div>
                <div className="grid col-span-1 pb-20 items-end ">
                    <Link
                        href="/api/auth/logout"
                        passHref
                        legacyBehavior
                    >
                        <LinkButton 
                            onClick={()=>{}}
                            buttonType="logoutButton"
                        >
                            Logout
                        </LinkButton>
                    </Link>
                    
                </div>
                
               
                   
            
            </div>
        </div>
        <div className="col-span-4 bg-[#171720]">{
            pathChangeLoad ? <div
                className="grid w-full h-full items-center justify-center"
            >
                <p className="text-blue-700 font-semibold" >Loading...</p>
            </div> :
            children
        }</div>
    </div>
  )
}

export default Dashboard
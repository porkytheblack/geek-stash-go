import '@/styles/globals.css'
import type { AppProps } from 'next/app'
import { UserProvider } from '@auth0/nextjs-auth0/client'
import Layout from '@/layout/main'

export default function App({ Component, pageProps }: AppProps) {
  return (
    <UserProvider>
      <Layout pageProps={pageProps} >
        <Component {...(pageProps)} />
      </Layout>
    </UserProvider>
  )
}

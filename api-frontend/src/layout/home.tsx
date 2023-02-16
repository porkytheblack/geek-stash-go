import React from 'react'

type tProps = {
    children: React.ReactNode
}

function Home( props: tProps ) {

    const { children } = props
  return (
    <div className="grid items-center justify-center">
        {
            children
        }
    </div>
  )
}

export default Home
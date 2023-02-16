import Image from 'next/image'
import React from 'react'

type tProps = {
    children: React.ReactNode
}

function Home( props: tProps ) {

    const { children } = props
  return (
    <div className="grid items-center justify-center">
        <Image
          src={`/photo.jpeg`}
          alt="Picture of the author"
          width={500}
          height={500}
        />
        {
            children
        }
    </div>
  )
}

export default Home
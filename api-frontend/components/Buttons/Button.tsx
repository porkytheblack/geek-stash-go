import React from 'react'
import twButtons from './twButtonStyles'

interface IProps {
    children: React.ReactNode,
    buttonType: keyof typeof twButtons,
    onClick?: (ev: any) => void
}

function Button(props: IProps) {
    const {buttonType, onClick, children} = props
  return (
    <button
        onClick={onClick}
        className={twButtons[buttonType]}
    >
        {children}
    </button>
  )
}

export default Button
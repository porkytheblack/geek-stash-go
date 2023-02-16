import React, { ForwardedRef, forwardRef } from 'react'
import twButtons from './twButtonStyles'

interface IProps {
    children: React.ReactNode,
    href?: string,
    buttonType: keyof typeof twButtons
    onClick?: (e: React.MouseEvent<HTMLAnchorElement>) => void
}

// eslint-disable-next-line react/display-name
const LinkButton = forwardRef((props: IProps, ref: ForwardedRef<HTMLAnchorElement>) => {

    const { children, href, buttonType, onClick } = props
    return (
    <a
        ref={ref}
        href={href}
        onClick={onClick}
        className={twButtons[buttonType]}
    >{children}</a>
    )
})

export default LinkButton
import './Button.css'


function Button(props) {
    var className = "Home-button"

    if ( props.opacity ) {
        className += " Home-button-opacity";
    }

    return (
        <div class={className}>
            {props.children}
        </div>
    )
}

export default Button;
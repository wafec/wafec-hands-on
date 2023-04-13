import './Background.css'


function Background(props) {
    return (
        <div class="Home-background">
            {props.children}
        </div>
    )
}

export default Background;
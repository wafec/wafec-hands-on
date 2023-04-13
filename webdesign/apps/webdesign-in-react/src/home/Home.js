import Background from "./Background";
import Image from "./Image";
import './Home.css'
import Title from "./Title";
import Button from "./Button";


function Home() {
    return (
        <Background>
            <div class="Home-container Home-flex">
                <div class="Home-flexitem-1">
                    <div>
                        <Title />
                    </div>
                    <div>
                        <Button opacity={false}><strong>Buy Now</strong></Button>
                        <Button opacity={true}>Shop All</Button>
                    </div>
                </div>
                <div class="Home-flexitem-1">
                    <Image />
                </div>
            </div>
        </Background>
    )
}

export default Home;
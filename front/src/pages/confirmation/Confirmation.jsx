import { Link } from "react-router-dom";
const Confirmation = () => {
  return (
    <div>
        <h1>Vous avez reçut un mail de confirmation dans votre boite mail</h1>
        <Link to="/">Retour à l'accueil</Link>
    </div>
  )
}

export default Confirmation
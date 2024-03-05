import { useEffect,useState } from 'react'
import { useParams } from 'react-router-dom';
import { MapPin,PhoneCall } from 'lucide-react';
import { Link } from 'react-router-dom';
import './salon.css';
const Salon = () => {
    const [salon, setSalon] = useState([]);
    const getDayOfWeekStyle = (dayOfWeek) => {
        const currentDay = new Date().getUTCDay();
        const days = ['Dimanche','Lundi', 'Mardi', 'Mercredi', 'Jeudi', 'Vendredi', 'Samedi'];
        return dayOfWeek === days[currentDay] ? { fontWeight: 'bold' } : {};
      };
    let { id } = useParams();
    useEffect(() => {
        document.title = 'Coiffeur: Les meilleurs coiffeurs à Paris 75000'
        fetch(`http://localhost:8097/salons/${id}`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json'
        }
        })
        .then(data => data.json())
        .then(data => {
            setSalon(data)
        })
        
    }
    , [id]);

    return (
        <div className="salonPage">
            <header>
                <div className="info">
                    <h2>{salon.Name}</h2>
                    <div className="coordonnes">
                        <p><MapPin />{salon.Address}</p>
                        <p><PhoneCall />{salon.Phone}</p>
                    </div>
                </div>
                <a href="#prestation" className="btn-primary" >Prendre RDV</a>
            </header>
            <div className="parent">
                <div className="div1"></div>
                <div className="div2"></div>
                <div className="div3"></div>
                <div className="div4"></div>
                <div className="div5"></div>
            </div>
            <p className="description description-salon">{salon.Description}</p>
            <h2>Réserver en ligne pour un RDV chez {salon.Name}</h2>
            <p className="description">24h/24 - Paiement sur place - Confirmation par mail</p>
            <section className="prestation" id="prestation">
                <h3>Choix de la prestation</h3>
                <div className="flex-container">
                    <div className="prestation-content">
                        {salon.Service && salon.Service.map((service, index) => (
                            <div key={index}>
                                <h4>{service.Name}</h4>
                                <p className="description">{service.Description}</p>
                                <div className="service-container">
                                {service.Prestation && service.Prestation.map((prestation, index) => (
                                    <div className="service-content" key={index}>
                                        <h5>{prestation.Name}</h5>
                                        <p className="price">{prestation.Price}€</p>
                                        <Link to={`/salons/${salon.ID}/reservation`} state={{ salonsData: salon, prestationData: prestation }} className='btn-primary' >Choisir</Link>
                                    </div>
                                ))}
                                </div>
                            </div>
                        ))}
                    </div>
                    <div className="salon-info">
                        <div className="rate-container">
                            <h4>Note Global</h4>
                            <div className="rate-content">
                                <div className="left">
                                   <span>4.9</span>
                                </div>
                                <div className="right">
                                    <ul>
                                        <li>Accueil <strong>4.9</strong></li>
                                        <li>Propreté <strong>4.9</strong></li>
                                        <li>Cadre & Ambiance <strong>4.9</strong></li>
                                        <li>Qualité de la prestation <strong>4.9</strong></li>
                                    </ul>
                                    <p>1006 clients ont donné leur avis</p>
                                </div>
                            </div>
                        </div>
                        <h4>Horaires d'ouverture</h4>
                        <div className="hourly-container">
                        {salon.Hours && salon.Hours.map((hour, index) => (
                        <p className="hourly" key={index} style={getDayOfWeekStyle(hour.DayOfWeek)}>
                            {hour.DayOfWeek} 
                            <span>
                            {new Date(hour.OpeningTime).getUTCHours().toString().padStart(2, '0') + ':' + new Date(hour.OpeningTime).getUTCMinutes().toString().padStart(2, '0')} - {new Date(hour.ClosingTime).getUTCHours().toString().padStart(2, '0') + ':' + new Date(hour.ClosingTime).getUTCMinutes().toString().padStart(2, '0')}
                            </span>
                        </p>
                        ))}
                        </div>
                    </div>
                </div>
            </section>
        </div>
    )
}

export default Salon
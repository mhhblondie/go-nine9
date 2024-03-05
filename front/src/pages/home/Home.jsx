import { useEffect,useState } from 'react'
import { Link } from 'react-router-dom'
import { MapPin,PhoneCall } from 'lucide-react';
import './home.css';

const Home = () => {
  const [salons, setSalons] = useState([]);
  useEffect(() => {
    document.title = 'Coiffeur: Les meilleurs coiffeurs à Paris 75000'
    fetch('http://localhost:8097/salons', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json'
      }
    })
    .then(data => data.json())
    .then(data => {
      setSalons(data)
    })
  }, []);

  return (
    <main className='homePage'>
      <div className='left'>
        <div className="intro">
          <h1>Sélectionnez un salon</h1>
          <p>Les meilleurs salon et instituts aux alentours de vous :  Réservation en ligne</p>
        </div>
        {salons.map((salon) => (
          <div key={salon.ID} className='salons'>
            <img src="https://res.cloudinary.com/planity/image/upload/x_0,y_216,w_2992,h_1780,c_crop/t_d_search,f_auto/m2mfnjksc00bdynoj9ag" alt="" />
            <div className="salon-info">
              <h2>{salon.Name}</h2>
              <p className='description'>{salon.Description.substring(0, 100)}...</p>
              <p><MapPin />{salon.Address}</p>
              <p><PhoneCall />{salon.Phone}</p>
            </div>
            <div className="cta-container">
              <Link to={`/salons/${salon.ID}`} className='btn-primary'>Prendre RDV</Link>
            </div>
          </div>
        ))}
      </div>
      <div className="right">
        <img src="/Map.PNG" alt="" />
      </div>
    </main>
  )
}

export default Home
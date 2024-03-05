import { useState, useEffect } from 'react';
import { useLocation } from 'react-router-dom'
import { MapPin, PhoneCall,CircleUser } from 'lucide-react';
import { Link,useNavigate } from 'react-router-dom';
import Cookies from 'js-cookie'

import './reservation.css';
const Reservation = () => {
    const location = useLocation()
    const [salon, setSalon] = useState({});
    const [prestation, setPrestation] = useState({});
    const [selectedUser, setSelectedUser] = useState();
    const [dategroupe,setDategroupe] = useState([]);
    const [isConnected, setIsConnected] = useState(false);
    const [selectedSlot, setSelectedSlot] = useState();
    const navigate = useNavigate();
    const groupedDates = dategroupe.reduce((groups, slot) => {
        const day = slot.time.split('T')[0];
        if (!groups[day]) {
          groups[day] = [];
        }
        groups[day].push(slot);
        return groups;
    }, {});
    
    const handleSlotClick = (slot) => {
        setSelectedSlot(slot);
    };
    
    const handleInputChange = (event) => {
        setSelectedUser(event.target.value);
        setDategroupe(salon.User.filter(staff => staff.ID === event.target.value)[0].Slots.map(slot => ({ time: slot.SlotTime, id: slot.ID })));
      };

    const handleSubmit = (event) => {
        event.preventDefault();
        let jwtToken = Cookies.get('authToken');
        fetch('http://localhost:8097/api/reservations', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization' : "Bearer "+ jwtToken
            },
            body: JSON.stringify({
                SlotID: selectedSlot,
            })
        })
        .then(response => {
            if (response.ok) {
                navigate('/confirmation');
            } else {
                console.error('Failed to make reservation');
            }
        });
    }
    
    useEffect(() => {
        if (location.state) {
            const { salonsData, prestationData } = location.state;
            setSalon(salonsData);
            setPrestation(prestationData);
            setSelectedUser(salonsData.User[0].ID);
            setDategroupe(salonsData.User[0].Slots.map(slot => ({ time: slot.SlotTime, id: slot.ID })));
            if(Cookies.get('authToken')) {
                setIsConnected(true);
            }else{
                setIsConnected(false);
            }
        }
      }, [location.state]);

  return (
    <div className='reservationPage'>
        <div className="info">
            <h1>{salon.Name}</h1>
            <div className="coordonnes">
                <p><MapPin />{salon.Address}</p>
                <p><PhoneCall />{salon.Phone}</p>
            </div>
        </div>
        <div className="first-step">
            <h2><span className="purple-bold">1.</span> Prestation sélecitonnée</h2>
            <div className="prestation-container">
                <div className="top">
                    <div className="information">
                        <h3>{prestation.Name}</h3>
                        <p>{prestation.Price}€</p>
                    </div>
                    <Link to={`/salons/${salon.ID}`} className="back">Supprimer</Link>
                </div>
                <div className="bottom">
                    <p>Avec qui ?</p>
                    <ul>
                        {salon.User && salon.User.map((user, index) => (
                            <li key={index}>
                                <label >
                                    <div className="selection">
                                        <CircleUser size={32}/>
                                        {user.Firstname}
                                    </div>
                                    <input type="radio" name="user" value={user.ID} defaultChecked={index === 0} onChange={handleInputChange} />
                                </label>
                            </li>
                        ))}
                    </ul>
                </div>
            </div>
        </div>
        <div className="second-step">
            <h2><span className='purple-bold'>2.</span> Choix de la date & heure</h2>
            <div className="hourly-container">
                {salon.User && salon.User.filter(staff => staff.ID === selectedUser).map((staff, index) => (
                    <div key={index}>
                        <div className="planning-container">
                            {Object.entries(groupedDates).map(([day, times], index) => (
                                <div key={index} className='day-container'>
                                <div className='day'>{new Date(day).toLocaleDateString('fr-FR', { weekday: 'long', day: 'numeric', month: 'long' })}</div>
                                {times.map((slot, index) => (
                                  <button
                                    className={`hour ${slot.id === selectedSlot ? 'selected' : ''}`}
                                    key={index}
                                    onClick={() => handleSlotClick(slot.id)}
                                >
                                  {new Date(slot.time).getUTCHours().toString().padStart(2, '0') + ':' + new Date(slot.time).getUTCMinutes().toString().padStart(2, '0')}
                                </button>
                                ))}
                              </div>
                            ))}
                        </div>
                    </div>
                ))}
            </div>
        </div>
        <div className="third-step">
        {isConnected ? (
            <>
                <h2><span className="purple-bold">3.</span> Réservation</h2>
                <button className="reservation btn-primary" onClick={handleSubmit}>Réserver</button>
            </>
        ) : (
            <>
                <h2><span className="purple-bold">3.</span> Identification</h2>
                <div className="account-container">
                    <h4>Nouveau sur GoNine ?</h4>
                    <Link to='/register' className="btn-secondary">Créer mon compte</Link>
                    <span className="separation">ou</span>
                    <h4>Vous avez déjà utilisé GoNine?</h4>
                    <Link to='/login' className="btn-primary">Se connecter</Link>
                </div>

            </>
        )}
        </div>
    </div>
  )
}

export default Reservation
import React, {useState} from 'react'
import { getCookie, setCookie } from '../../AuthContext/AuthContext';
import { useNavigate, useLocation } from 'react-router-dom';


const ModifyMySalon = () => {
    const navigate = useNavigate();
    let { state  } = useLocation();
    console.log(state.state)

    const [ModifiedSalon, setModifiedSalon] = useState(state.state);
    const [errors, setErrors] = useState("")

    const token = getCookie("authToken")



    // Change les valeurs du salon
    const handleInputChange = (event) => {
        const { name, value } = event.target;
        setModifiedSalon({ ...ModifiedSalon, [name]: value });
      };

    const handleModify = async() =>{
        const {Name, Address, Phone, user, Description} = ModifiedSalon;
        const requestData = {Name, Address, Phone, user, Description};
        try {
            const response = await fetch(`http://localhost:8097/api/management/salons/${state.state.ID}`, {
              method: 'PATCH',
              headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${token}`,
              },
              mode: 'cors',
              body: JSON.stringify(requestData),
            });
        
            if (!response.ok) {
              const errorResponse = await response.json();
              console.log(errorResponse)
              setErrors(errorResponse.message || 'Échec de la requête d\'inscription');
            }
            if(response.ok){
                setErrors("")
                navigate("/admin/")
            }
           
            

          } catch (err) {
            setErrors(err.message || 'Une erreur inattendue s\'est produite');
          }

    }


      
  return (
    <div>
        <h1> Modifier un Salon </h1>
        {errors.length > 0 && <h3>{errors}</h3>}
        <div>
        {ModifiedSalon && 
        
            <>
            <button onClick={handleModify}>Modifies ton salon</button>
            <label>Nom du salon</label>
            <input name="Name" type="text" placeholder='Ex: Mon salon' value={ModifiedSalon.Name} onChange={handleInputChange}/>
            <label>Adresse</label>
            <input name="Address" type="text" placeholder='Ex: 10 rue de Louvois 75002 Paris' value={ModifiedSalon.Address} onChange={handleInputChange}/>
            <label>Numéro de téléphone</label>
            <input name="Phone" type="tel" placeholder="Ex: 667" value={ModifiedSalon.Phone} onChange={handleInputChange}/>
            <label>Description</label>
            <input name="Description" type="textarea" placeholder="Ex: Description..." value={ModifiedSalon.Description} onChange={handleInputChange}/>
  
            
          
            </>
            }

            </div>
      
    </div>
  )
}

export default ModifyMySalon

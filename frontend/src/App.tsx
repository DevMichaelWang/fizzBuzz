import { useEffect, useState} from 'react';
import axios from 'axios';
import './App.css';

const App = () => {

  const [counter, setCounter] = useState<number>(1);
  const [message, setMessage] = useState<string>('');
  const API_URL: string | undefined = process.env.REACT_APP_API_URL;

  const getData = async(count: number) => {       
    try {      
      const res = await axios.post(`${API_URL}/fizzbuzz`, { count: count })      
      if (res.data.status === 200) setMessage(res.data.message);
    } catch (e) {
      console.log ("error", e)      
    }    
  }

  useEffect (()=> {    
    getData (counter)
  }, [counter])

  const onClickButton = () => {    
    setCounter (counter + 1)        
  }  

  return (
    <div className='app-container'>
      <div className='align-center'>
        <div className='counter'>Your count</div>
        <div className='counter'>{counter}</div>
        <button onClick={onClickButton}>Push me!</button>
        <div className='display-result'>{message}</div>
      </div>
    </div>
  );

}

export default App;

import './App.css';
import axios from 'axios'
import React, { useEffect, useState } from 'react';

function App() {

  const [batteryData, setBatteryData] = useState([]);

  useEffect(() => {
    axios.get('http://localhost:3001/getBatteries')
      .then(response => {
        setBatteryData(response.data.Data)
      })
      .catch(error => console.log(error));
  }, []);

  return (
    <table>
      <thead>
        <tr>
          <th>ID</th>
          <th>Current Health</th>
          <th>Maximum Capacityyyy</th>
        </tr>
      </thead>
      <tbody>
        {batteryData.map(battery => (
          <tr key={battery.id}>
            <td>{battery.id}</td>
            <td>{battery.current_health}</td>
            <td>{battery.maximum_capacity}</td>
          </tr>
        ))}
      </tbody>
    </table>
  );
}

export default App;

import './App.css';
import { Amplify } from 'aws-amplify';
import config from './amplifyconfiguration.json';
import { signIn, fetchAuthSession, getCurrentUser, signOut } from 'aws-amplify/auth';
import { useEffect, useState } from 'react';
import { TextAreaField } from '@aws-amplify/ui-react';

Amplify.configure(config)

function App() {
  const [email, setEmail] = useState('pjmessi25@icloud.com');
  const [password, setPassword] = useState('Password123!');
  const [user, setUser] = useState(null)
  const [authToken, setAuthToken] = useState('')

  useEffect(() => {
    getCurrentUser().then((userDetails) => {
      setUser(userDetails)

      fetchAuthSession().then(sessionDetails => {
        setAuthToken(sessionDetails.tokens?.accessToken?.toString())
      })
    }).catch(e => {
      console.log('user not authenticated: ', e)
    });
  }, [])

  const signInUser = async () => {
    await signIn({
      username: email,
      password: password,
    })

    const userInfo = await getCurrentUser()
    setUser(userInfo)
    const sessionDetails = await fetchAuthSession()
    setAuthToken(sessionDetails.tokens?.accessToken?.toString())
  }

  const signOutUser = async () => {
    await signOut()
    setUser(null)
    setAuthToken('')
  }

  const copyToClipboard = () => {
    const text = authToken;
    if (!navigator.clipboard) {
      // Fallback for older browsers (pre-Clipboard API)
      const textArea = document.createElement('textarea');
      textArea.value = text;
      document.body.appendChild(textArea);
      textArea.select();
      try {
        document.execCommand('copy'); // Try to copy text
        console.log('Text copied to clipboard (fallback)');
      } catch (err) {
        console.error('Failed to copy text to clipboard (fallback)', err);
      }
      document.body.removeChild(textArea);
    } else {
      // Modern (Clipboard API) approach
      navigator.clipboard.writeText(text)
        .then(() => {
          console.log('Text copied to clipboard (Clipboard API)');
        })
        .catch((err) => {
          console.error('Failed to copy text to clipboard (Clipboard API)', err);
        });
    }
  }

  return (
    <div className="App">
      <div style={{ marginTop: "1%" }}>
        <input placeholder="email" value={email} onChange={(e) => setEmail(e.target.value)} type="text" /><br />
        <input placeholder="password" type="text" value={password} onChange={(e) => setPassword(e.target.value)} />
        <br />
        {
          !Boolean(user) ? <button onClick={signInUser}>Sign In</button> : <button onClick={signOutUser}>Sign Out</button>
        }
      </div>

      {authToken && (<>
        <TextAreaField disabled style={{ marginTop: "1%" }} rows="10" cols="50" defaultValue={authToken} />
        <button onClick={copyToClipboard}>copy to clipboard</button>
      </>)}
    </div>
  );
}

export default App;

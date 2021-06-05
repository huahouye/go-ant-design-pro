import React, { useState } from 'react';
import styles from './index.css';
import zxcvbn from 'zxcvbn';

export default function Page() {

    const [password, setPassword] = useState('')

    return (
        <div>
            <input autoComplete="off" type="password" onChange={e => setPassword(e.target.value)} />
            <h1 className={styles.title}>{password}</h1>
            <progress className={`password-strength-meter-progress strength-${zxcvbn(password).score}`}
                value={zxcvbn(password).score}
                max="4"
            />
            {password && (
                <div>
                    <strong>Password strength score:</strong> {zxcvbn(password).score}
                </div>
            )}
        </div>
    )
}



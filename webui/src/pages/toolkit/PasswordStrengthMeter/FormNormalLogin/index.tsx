import React, { useState } from "react";
import styles from "./index.less";
import { Form, Input, Button, Checkbox, Progress } from "antd";
import { UserOutlined, LockOutlined } from "@ant-design/icons";
import zxcvbn from 'zxcvbn';

const PasswordStrengthMeter = (props: any) => {

  const { password } = props;
  const testedResult = zxcvbn(password);

  const passwordLevel = (result: any) => {
    switch (result.score) {
      case 0:
        return '很弱';
      case 1:
        return '较弱';
      case 2:
        return '一般';
      case 3:
        return '很好';
      case 4:
        return '很强';
      default:
        return '很弱';
    }
  }

  return (
    <div>
      <Progress 
        showInfo={false} 
        strokeColor={{
          '0%': '#108ee9',
          '100%': '#87d068',
        }}
        percent={(testedResult.score/4)*100} />
      <br />
      <label>
        {password && (
          <>
            <strong>密码强度:</strong> {passwordLevel(testedResult)}
          </>
        )}
      </label>
    </div>
  );
}

const NormalLoginForm = () => {

  const onFinish = (values: any) => {
    console.log("Received values of form: ", values);
  };

  const [password, setPassword] = useState("")

  return (
    <Form
      name="normal_login"
      className="login-form"
      initialValues={{ remember: true }}
      onFinish={onFinish}
    >
      <Form.Item
        name="username"
        rules={[{ required: true, message: "Please input your Username!" }]}
      >
        <Input
          prefix={<UserOutlined className="site-form-item-icon" />}
          placeholder="Username"
        />
      </Form.Item>
      <Form.Item
        name="password"
        rules={[{ required: true, message: "Please input your Password!" }]}
      >
        <Input
          prefix={<LockOutlined className="site-form-item-icon" />}
          type="password"
          placeholder="Password"
          onChange={(e) => setPassword(e.target.value)}
        />
        <PasswordStrengthMeter password={password}/>
      </Form.Item>
      <Form.Item>
        <Form.Item name="remember" valuePropName="checked" noStyle>
          <Checkbox>Remember me</Checkbox>
        </Form.Item>

        <a className="login-form-forgot" href="">
          Forgot password
        </a>
      </Form.Item>

      <Form.Item>
        <Button type="primary" htmlType="submit" className="login-form-button">
          Log in
        </Button>
        Or <a href="">register now!</a>
      </Form.Item>
    </Form>
  );
};

export default () => (
  <div className={styles.container}>
    <div id="components-form-demo-normal-login">
      <NormalLoginForm />
    </div>
  </div>
);

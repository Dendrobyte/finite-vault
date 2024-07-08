import { Body, Controller, Post } from '@nestjs/common';
import axios from 'axios';
import { OAuth2Client } from 'google-auth-library';
import { AppService } from './app.service';

// TODO: Worth moving this out so I can error catch here
const client = new OAuth2Client(
  process.env.GOOGLE_CLIENT_ID,
  process.env.GOOGLE_CLIENT_SECRET,
);

// TODO: Is this the right "dependency injection way"?
//       https://docs.nestjs.com/techniques/http-module
// TODO: Take an obejct, not two variables
const verifySimpleLoginCode = async (code, redirect_uri) => {
  // https://simplelogin.io/docs/siwsl/code-flow/
  const simpleLoginData = {
    grant_type: 'authorization_code',
    code: code,
    redirect_uri: redirect_uri,
    client_id: process.env.SIMPLE_LOGIN_CLIENT_ID,
    client_secret: process.env.SIMPLE_LOGIN_CLIENT_SECRET,
  };

  const allInfo = await axios
    .post('https://app.simplelogin.io/oauth2/token', simpleLoginData, {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    })
    .then((resp) => {
      return resp.data;
    })
    .catch((err) => {
      console.log('Error on SimpleLogin OAuth: ' + err);
      console.log(err.response);
      return undefined;
    });

  return allInfo.user;
};

@Controller()
export class AppController {
  constructor(private readonly appService: AppService) {}

  // TODO: Is there a better way to organize this?
  //       Make an if statement that changes the process based on a service name, thus we aren't making duplicate calls to appService.loginUser
  @Post('/login_google')
  async loginGoogle(@Body('token') token): Promise<any> {
    const ticket = await client.verifyIdToken({
      idToken: token,
      audience: process.env.GOOGLE_CLIENT_ID,
    });

    const { name, email } = ticket.getPayload(); // Deconstruct necessary fields from the payload
    const user_info = await this.appService.loginUser({ name, email });
    return {
      user_info,
      message: 'Success',
    };
  }

  @Post('/login_proton')
  async loginProton(
    @Body('token') token: string,
    @Body('redirect_uri') redirect_uri: string,
  ): Promise<any> {
    const { name, email } = await verifySimpleLoginCode(token, redirect_uri);
    const user_info = await this.appService.loginUser({ name, email });
    return {
      user_info,
      message: 'Success',
    };
  }
}

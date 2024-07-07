import { Body, Controller, Post } from '@nestjs/common';
import { OAuth2Client } from 'google-auth-library';
import { AppService } from './app.service';

const client = new OAuth2Client(
  process.env.GOOGLE_CLIENT_ID,
  process.env.GOOGLE_CLIENT_SECRET,
);

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
  async loginProton(@Body('token') token): Promise<any> {
    console.log('User auth token accepted is: ' + token);
  }
}

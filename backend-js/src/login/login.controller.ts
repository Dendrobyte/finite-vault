import {
  Body,
  Controller,
  Get,
  HttpStatus,
  Param,
  Post,
  Res,
} from '@nestjs/common';
import { Response } from 'express';
import { LoginService } from './login.service';

type UserLoginResponseObject = {
  username: string;
  email: string;
  authToken: string; // TODO: This might become a JWT type or something, but for now let's just string.
};

@Controller('login')
export class LoginController {
  constructor(private readonly loginService: LoginService) {}

  @Get()
  default(): string {
    return 'Ohaiio';
  }

  /* THIS IS A DEVELOPMENT LOGIN FUNCTION TO AVOID REQUIRING FRONTEND. REMOVE ON PRODUCTION. */
  // TODO: Create a token generation tool if necessary, otherwise use a built-up test suite
  @Post('test')
  async fakeLogin() {
    const token = this.loginService.generateJWT({ email: 'test@test.com' });
  }

  @Post(':service')
  async mainLogin(
    @Param('service') service: string,
    @Body('token') token: string,
    @Body('redirect_uri') redirect_uri: string,
    @Res({ passthrough: true }) res: Response,
  ) {
    if (service === 'google') {
      console.log('Triggering login with google');
      await this.loginService
        .loginGoogle(token)
        .then((userInfo) => {
          const user_data: UserLoginResponseObject = {
            username: userInfo.username,
            email: userInfo.email,
            authToken: 'abc123',
          };
          return {
            user_data,
            message: 'Success',
          };
        })
        .catch((err) => {
          console.log(`Error logging in user with Google: ${err}`);
          res.status(HttpStatus.UNPROCESSABLE_ENTITY);
        });
    } else if (service === 'proton') {
      await this.loginService
        .loginProton(token, redirect_uri)
        .then((userInfo) => {
          console.log(userInfo);
          const user_data: UserLoginResponseObject = {
            username: userInfo.username,
            email: userInfo.email,
            authToken: 'xyz456',
          };
          return {
            user_data,
            message: 'Success',
          };
        })
        .catch((err) => {
          console.log(`Error logging in user with SimpleLogin: ${err}`);
          res.status(HttpStatus.UNPROCESSABLE_ENTITY);
        });
    } else {
      res.status(HttpStatus.NOT_FOUND);
      return 'illegal';
    }
  }
}

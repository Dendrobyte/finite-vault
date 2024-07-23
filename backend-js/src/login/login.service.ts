import { Injectable } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import axios from 'axios';
import { OAuth2Client } from 'google-auth-library';
import { Model } from 'mongoose';
import { User, UserDocument } from 'src/db/user.schema';
import { JwtService } from '@nestjs/jwt';

const googleClient = new OAuth2Client(
  process.env.GOOGLE_CLIENT_ID,
  process.env.GOOGLE_CLIENT_SECRET,
);

type UserInfo = {
  username: string;
  email: string;
  balance: number;
};

type JWTPayload = {
  email: string;
};

// Helper function to make relevant calls to simple login authorization
const verifySimpleLoginCode = async (token: string, redirect_uri: string) => {
  // https://simplelogin.io/docs/siwsl/code-flow/
  const simpleLoginData = {
    grant_type: 'authorization_code',
    code: token,
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

@Injectable()
export class LoginService {
  constructor(
    @InjectModel(User.name) private userModel: Model<UserDocument>,
    private jwtService: JwtService,
  ) {}

  // Create a JWT token for authentication purposes
  async generateJWT(payload: JWTPayload) {
    // We only need the email for the payload
    const token = await this.jwtService.signAsync(payload.email);
    return token;
  }

  /* Hits the database to either create a user or find an existing one based on email */
  // TODO: Return JWT token as well here
  async loginUser({ name, email }): Promise<any> {
    const user = await this.userModel.findOne({ email: email });
    if (!user) {
      const newUser = new this.userModel({ email, name, balance: 0 });
      await newUser.save();
      return {
        username: newUser.name,
        email: newUser.email,
        balance: newUser.balance,
      };
    } else {
      return {
        username: user.name,
        email: user.email,
        balance: user.balance,
      };
    }
  }

  // TODO: The below functions should call the same JWT generation function
  // NOTE: Wrap the user's email in the JWT and ensure you only have access to that which is logged in
  //       I can "worry" about mismatches later

  /* Handle token verification with the Google client */
  async loginGoogle(token: any): Promise<any> {
    const ticket = await googleClient.verifyIdToken({
      idToken: token,
      audience: process.env.GOOGLE_CLIENT_ID,
    });

    const { name, email } = ticket.getPayload(); // Deconstruct necessary fields from the payload
    const userInfo = await this.loginUser({ name, email });
    return userInfo;
  }

  /* Handle login verification via Simple Login */
  async loginProton(token: string, redirect_uri: string) {
    const { name, email } = await verifySimpleLoginCode(token, redirect_uri);
    const userInfo: UserInfo = await this.loginUser({ name, email });
    return userInfo;
  }
}

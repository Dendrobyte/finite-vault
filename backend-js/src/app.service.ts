import { Injectable } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { Model } from 'mongoose';
import { User, UserDocument } from './db/user.schema';

// TODO: Move this to somewhere that makes more sense
type BaseUser = {
  email: string;
  name: string;
};

@Injectable()
export class AppService {
  constructor(@InjectModel(User.name) private userModel: Model<UserDocument>) {}

  // Handle obtaining initial information for a stored user in our Mongo database
  // TODO: Check for username field, if not present then use name?
  //       Or on initial creation, set username to name.
  async loginUser({ email, name }: BaseUser): Promise<any> {
    const user = await this.userModel.findOne({ email: email });
    if (!user) {
      const newUser = new this.userModel({ email, name, balance: 0 });
      await newUser.save();
      return newUser;
    } else {
      return user;
    }
  }
}

import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';

export type UserDocument = User & Document;

@Schema()
export class User {
  @Prop()
  name: string;

  @Prop()
  email: string;

  @Prop()
  balance: number;

  // TODO: Other fields for users
}

export const UserSchema = SchemaFactory.createForClass(User);

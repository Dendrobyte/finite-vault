import { HttpModule } from '@nestjs/axios';
import { Module } from '@nestjs/common';
import { ConfigModule } from '@nestjs/config';
import { MongooseModule } from '@nestjs/mongoose';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { User, UserSchema } from './db/user.schema';
import { LoginController } from './login/login.controller';
import { LoginService } from './login/login.service';

// TODO: BIG! Please refactor prior to doign other routes so this becomes the LoginModule
@Module({
  imports: [
    ConfigModule.forRoot({
      isGlobal: true,
      envFilePath: '.env',
    }),
    MongooseModule.forRoot(process.env.MONGO_URI),
    MongooseModule.forFeature([{ name: User.name, schema: UserSchema }]),
    HttpModule,
  ],
  controllers: [AppController, LoginController],
  providers: [AppService, LoginService],
})
export class AppModule {}

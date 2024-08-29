import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);

  // Hello, my old friend (CORS)
  app.enableCors({
    origin: 'http://localhost:5173', // TODO: Add the cloudflare URL
  });

  await app.listen(8080);
  console.log('Finite Vault backend is listening!');
}
bootstrap();

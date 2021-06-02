import { SwaggerModule, DocumentBuilder } from '@nestjs/swagger';

export class BaseAPIDocumnetation {
  public builder = new DocumentBuilder();

  public initializeOptions() {
    return this.builder
      .setTitle('NFC API SERVER')
      .setDescription('BACKEND API')
      .setVersion('1.0')
      .build();
  }
}

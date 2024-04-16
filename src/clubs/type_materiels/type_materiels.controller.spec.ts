import { Test, TestingModule } from '@nestjs/testing';
import { TypeMaterielsController } from './type_materiels.controller';
import { TypeMaterielsService } from './type_materiels.service';

describe('TypeMaterielsController', () => {
  let controller: TypeMaterielsController;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      controllers: [TypeMaterielsController],
      providers: [TypeMaterielsService],
    }).compile();

    controller = module.get<TypeMaterielsController>(TypeMaterielsController);
  });

  it('should be defined', () => {
    expect(controller).toBeDefined();
  });
});

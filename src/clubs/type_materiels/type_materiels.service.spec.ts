import { Test, TestingModule } from '@nestjs/testing';
import { TypeMaterielsService } from './type_materiels.service';

describe('TypeMaterielsService', () => {
  let service: TypeMaterielsService;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      providers: [TypeMaterielsService],
    }).compile();

    service = module.get<TypeMaterielsService>(TypeMaterielsService);
  });

  it('should be defined', () => {
    expect(service).toBeDefined();
  });
});

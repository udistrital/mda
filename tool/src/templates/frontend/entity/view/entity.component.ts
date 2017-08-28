import {Component, OnInit} from '@angular/core';
import {JornadaService} from '../../services/jornada.service';
import {Jornada} from '../../models/jornada';
import {Router} from '@angular/router';
import {GlobalsComponent} from '../../globals/globals.component';
import {ConfirmationService} from 'primeng/primeng';

@Component({
  selector: 'app-jornada',
  templateUrl: './jornada.component.html',
  styleUrls: ['./jornada.component.css']
})
export class JornadaComponent implements OnInit {

  jornadas: Jornada[];
  jornada: Jornada;

  constructor(private jornadaService: JornadaService,
              private router: Router, private globals: GlobalsComponent,
              private confirmationService: ConfirmationService) {
    this.globals = globals;
  }

  ngOnInit(): void {
    this.jornadaService.getJornadas().then(jornadas => this.jornadas = jornadas);
  }

  newJornada(): void {

    this.router.navigate(['/jornada/new']).then(() => null);
    this.globals.currentModule = 'Jornada';
  }

  editar(jornada: Jornada): void {
    this.jornada = jornada;
    this.router.navigate(['/jornada/edit', this.jornada.id ]);
  }

  borrar(jornada: Jornada): void {
    this.confirmationService.confirm({
      message: 'Esta seguro que quiere borrar la jornada ' + jornada.nombre + '?',
      accept: () => {
        this.jornadaService.delete(jornada.id)
          .then(response => this.jornadaService.getJornadas().then(jornadas => this.jornadas = jornadas));
      }
    });
  }
}

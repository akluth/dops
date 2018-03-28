extern crate tui;
extern crate termion;

use std::io;
use std::thread;
use std::sync::mpsc;

use termion::event;
use termion::input::TermRead;

use tui::Terminal;
use tui::backend::RawBackend;
use tui::widgets::{Widget, Block, Borders};
//use tui::layout::{Group, Size, Direction};

enum Event {
    Input(event::Key),
}

fn main() {
    let mut terminal = init().expect("Failed initialization");

    let (tx, rx) = mpsc::channel();
    let input_tx = tx.clone();

    thread::spawn(move || {
        let stdin = io::stdin();
        for c in stdin.keys() {
            let evt = c.unwrap();
            input_tx.send(Event::Input(evt)).unwrap();
            if evt == event::Key::Char('q') {
                break;
            }
        }
    });

    loop {
        draw(&mut terminal).expect("Failed to draw");

        let evt = rx.recv().unwrap();
        match evt {
            Event::Input(input) => match input {
                event::Key::Char('q') => {
                    break;
                }
                _ => {}
            }
        }
    }

    terminal.show_cursor().unwrap();
    terminal.clear().unwrap();
}

fn init() -> Result<Terminal<RawBackend>, io::Error> {
    let backend = RawBackend::new()?;
    Terminal::new(backend)
}

fn draw(t: &mut Terminal<RawBackend>) -> Result<(), io::Error> {

    let size = t.size()?;

    Block::default()
        .title("Block")
        .borders(Borders::ALL)
        .render(t, &size);

    t.draw()
}

#![no_std]
use soroban_sdk::{contract, contractimpl, symbol_short, Env, Symbol};

const DISBURSED: Symbol = symbol_short!("disbursed");

#[contract]
pub struct ConditionalEscrow;

#[contractimpl]
impl ConditionalEscrow {
    pub fn release_payment(env: Env, amount: i128) -> i128 {
        // Emit typed event for the Go relay daemon
        env.events().publish((DISBURSED,), amount);
        amount
    }
}

#[cfg(test)]
mod test {
    use super::*;
    use soroban_sdk::Env;

    #[test]
    fn test_release_payment() {
        let env = Env::default();
        let contract_id = env.register(ConditionalEscrow, ());
        let client = ConditionalEscrowClient::new(&env, &contract_id);

        let payout = client.release_payment(&500);
        assert_eq!(payout, 500);
    }
}

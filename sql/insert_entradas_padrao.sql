﻿-- INSERT INTO tipo_conta (
-- 	nome, descricao_debito, descricao_credito)
-- VALUES (
-- 	'banco', 'saque', 'depósito');
-- 
-- INSERT INTO tipo_conta (
-- 	nome, descricao_debito, descricao_credito)
-- VALUES (
-- 	'carteira', 'gastar', 'receber')
-- RETURNING nome;  -- com retorno de valor
-- 
-- INSERT INTO tipo_conta (
-- 	nome, descricao_debito, descricao_credito)
-- VALUES (
-- 	'despesa', 'desconto', 'despesa'), (
-- 	'cartão de crédito', 'cobrar', 'pagamento'), (
-- 	'receita', 'receita', 'cobrar');
-- 
-- INSERT INTO tipo_conta (
-- 	nome, descricao_debito, descricao_credito)
-- VALUES (
-- 	'ativo', 'descrescer', 'aumentar'), (
-- 	'passivo', 'aumentar', 'descrescer'), (
-- 	'líquido', 'aumentar', 'descrescer');
	
-- -- SELECT * FROM tipo_conta;


-- INSERT INTO conta (
-- 	nome, tipo_conta_nome, codigo, conta_pai, comentario)
-- VALUES (
-- 	'ativos', 'ativo', '1', null, '');  
-- 
-- INSERT INTO conta (
-- 	nome, tipo_conta_nome, codigo, conta_pai, comentario)
-- VALUES (
-- 	'despesas', 'despesa', '2', null, ''), (
-- 	'líquidos', 'líquido', '3', null, ''), (
-- 	'passivos', 'passivo', '4', null, ''), (
-- 	'receitas', 'receita', '5', null, '');
-- 
-- INSERT INTO conta (
-- 	nome, tipo_conta_nome, codigo, conta_pai, comentario)
-- VALUES (
-- 	'conta corrente', 'banco', '6', 'ativos', ''), (
-- 	'conta poupança', 'banco', '7', 'ativos', ''), (
-- 	'dinheiro em carteira', 'carteira', '8', 'ativos', '');
-- 
-- INSERT INTO conta (
-- 	nome, tipo_conta_nome, codigo, conta_pai, comentario)
-- VALUES (
-- 	'nu conta', 'banco', '9', 'conta corrente', ''), (
-- 	'bradesco', 'banco', '10', 'conta poupança', ''), (
-- 	'caixa econômica', 'banco', '11', 'conta poupança', '');
-- 
-- INSERT INTO conta (
-- 	nome, tipo_conta_nome, codigo, conta_pai, comentario)
-- VALUES (
-- 	'internet', 'despesa', '12', 'despesas', ''), (
-- 	'telefone', 'despesa', '13', 'despesas', ''), (
-- 	'serviços', 'despesa', '14', 'despesas', ''), (
-- 	'eletricidade', 'despesa', '15', 'serviços', ''), (
-- 	'refeições fora', 'despesa', '16', 'despesas', ''), (
-- 	'computador', 'despesa', '17', 'despesas', '');

-- INSERT INTO conta (
-- 	nome, tipo_conta_nome, codigo, conta_pai, comentario)
-- VALUES (
-- 	'cartão de crédito', 'passivo', '18', 'passivos', ''), (
-- 	'nubank', 'passivo', '18', 'cartão de crédito', ''), (
-- 	'submarino', 'passivo', '19', 'cartão de crédito', ''), (
-- 	'salário', 'receita', '20', 'receitas', ''), (
-- 	'juros recebidos', 'receita', '21', 'receitas', '');

-- -- SELECT * FROM conta;



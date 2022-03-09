INSERT INTO public.tipo_conta(
	nome, descricao_debito, descricao_credito, data_criacao)
	VALUES (
		'banco', 'saque', 'depósito', now()), (
		'carteira', 'gastar', 'receber', now()), (
		'despesa', 'desconto', 'despesa', now()), (
		'cartão de crédito', 'cobrar', 'pagamento', now()), (
		'receita', 'receita', 'cobrar', now()), (
		'ativo', 'descrescer', 'aumentar', now()), (
		'passivo', 'aumentar', 'descrescer', now()), (
		'líquido', 'aumentar', 'descrescer', now());

--

INSERT INTO public.conta(
	nome, nome_tipo_conta, codigo, conta_pai, comentario, data_criacao)
	VALUES (
		'ativos', 'ativo', '1', null, null, now()), (
        'despesas', 'despesa', '2', null, null, now()), (
        'líquidos', 'líquido', '3', null, null, now()), (
        'passivos', 'passivo', '4', null, null, now()), (
        'receitas', 'receita', '5', null, null, now()), (
        'conta corrente', 'banco', '6', 'ativos', null, now()), (
        'conta poupança', 'banco', '7', 'ativos', null, now()), (
        'dinheiro em carteira', 'carteira', '8', 'ativos', null, now()), (
        'nu conta', 'banco', '9', 'conta corrente', null, now()), (
        'bradesco', 'banco', '10', 'conta poupança', null, now()), (
        'caixa econômica', 'banco', '11', 'conta poupança', null, now()), (
        'internet', 'despesa', '12', 'despesas', null, now()), (
        'telefone', 'despesa', '13', 'despesas', null, now()), (
        'serviços', 'despesa', '14', 'despesas', null, now()), (
        'eletricidade', 'despesa', '15', 'serviços', null, now()), (
        'refeições fora', 'despesa', '16', 'despesas', null, now()), (
        'computador', 'despesa', '17', 'despesas', null, now());